package integration_tests

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/icza/dyno"

	xiontypes "github.com/burnt-labs/xion/x/xion/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	ibctest "github.com/strangelove-ventures/interchaintest/v7"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/testreporter"
	"github.com/strangelove-ventures/interchaintest/v7/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func BuildXionChain(t *testing.T) (*cosmos.CosmosChain, context.Context){
	ctx := context.Background()

	var numFullNodes = 1
	var numValidators = 3

	// pulling image from env to foster local dev
	imageTag := os.Getenv("XION_IMAGE")
	imageTagComponents := strings.Split(imageTag, ":")

	// Chain factory
	cf := ibctest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*ibctest.ChainSpec{
		{
			Name:    imageTagComponents[0],
			Version: imageTagComponents[1],
			ChainConfig: ibc.ChainConfig{
				Images: []ibc.DockerImage{
					{
						Repository: imageTagComponents[0],
						Version:    imageTagComponents[1],
						UidGid:     "1025:1025",
					},
				},
				GasPrices:              "0.0uxion",
				GasAdjustment:          1.3,
				Type:                   "cosmos",
				ChainID:                "xion-1",
				Bin:                    "xiond",
				Bech32Prefix:           "xion",
				Denom:                  "uxion",
				TrustingPeriod:         "336h",
				ModifyGenesis:          modifyGenesisShortProposals(votingPeriod, maxDepositPeriod),
				UsingNewGenesisCommand: true,
			},
			NumValidators: &numValidators,
			NumFullNodes:  &numFullNodes,
		},
	})

	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)

	xion := chains[0].(*cosmos.CosmosChain)

	// Relayer Factory
	client, network := ibctest.DockerSetup(t)
	//relayer := ibctest.NewBuiltinRelayerFactory(ibc.CosmosRly, zaptest.NewLogger(t)).Build(
	//	t, client, network)

	// Prep Interchain
	const ibcPath = "xion-osmo-dungeon-test"
	ic := ibctest.NewInterchain().
		AddChain(xion)
	//AddRelayer(relayer, "relayer").
	//AddLink(ibctest.InterchainLink{
	//	Chain1:  xion,
	//	Chain2:  osmosis,
	//	Relayer: relayer,
	//	Path:    ibcPath,
	//})

	// Log location
	f, err := ibctest.CreateLogFile(fmt.Sprintf("%d.json", time.Now().Unix()))
	require.NoError(t, err)
	// Reporter/logs
	rep := testreporter.NewReporter(f)
	eRep := rep.RelayerExecReporter(t)

	// Build Interchain
	require.NoError(t, ic.Build(ctx, eRep, ibctest.InterchainBuildOptions{
		TestName:          t.Name(),
		Client:            client,
		NetworkID:         network,
		BlockDatabaseFile: ibctest.DefaultBlockDatabaseFilepath(),

		SkipPathCreation: false},
	),
	)
	return xion, ctx
}

func TestXionSendPlatformFee(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}

	t.Parallel()

	xion, ctx := BuildXionChain(t)
	
	// Create and Fund User Wallets
	t.Log("creating and funding user accounts")
	fundAmount := int64(10_000_000)
	users := ibctest.GetAndFundTestUsers(t, ctx, "default", fundAmount, xion)
	xionUser := users[0]
	t.Logf("created xion user %s", xionUser.FormattedAddress())

	xionUserBalInitial, err := xion.GetBalance(ctx, xionUser.FormattedAddress(), xion.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, fundAmount, xionUserBalInitial)

	// step 1: send a xion message with default (0%) platform fee
	recipientKeyName := "recipient-key"
	err = xion.CreateKey(ctx, recipientKeyName)
	require.NoError(t, err)
	receipientKeyAddressBytes, err := xion.GetAddress(ctx, recipientKeyName)
	require.NoError(t, err)
	recipientKeyAddress, err := types.Bech32ifyAddressBytes(xion.Config().Bech32Prefix, receipientKeyAddressBytes)
	require.NoError(t, err)

	_, err = xion.FullNodes[0].ExecTx(ctx,
		xionUser.KeyName(),
		"xion", "send", xionUser.KeyName(),
		"--chain-id", xion.Config().ChainID,
		recipientKeyAddress, fmt.Sprintf("%d%s", 100, xion.Config().Denom),
	)
	require.NoError(t, err)
	balance, err := xion.GetBalance(ctx, recipientKeyAddress, xion.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, uint64(100), uint64(balance))

	// step 2: update the platform percentage to 5%
	config := types.GetConfig()
	config.SetBech32PrefixForAccount("xion", "xionpub")

	setPlatformPercentageMsg := xiontypes.MsgSetPlatformPercentage{
		Authority:          authtypes.NewModuleAddress("gov").String(),
		PlatformPercentage: 500,
	}

	xion.Config().EncodingConfig.InterfaceRegistry.RegisterImplementations(
		(*types.Msg)(nil),
		&xiontypes.MsgSetPlatformPercentage{},
	)
	cdc := codec.NewProtoCodec(xion.Config().EncodingConfig.InterfaceRegistry)

	msg, err := cdc.MarshalInterfaceJSON(&setPlatformPercentageMsg)

	prop := cosmos.Proposal{
		Messages: []json.RawMessage{msg},
		Metadata: "",
		Deposit:  "100uxion",
		Title:    "Set platform percentage to 5%",
		Summary:  "Ups the platform fee to 5% for the integration test",
	}
	paramChangeTx, err := xion.SubmitProposal(ctx, xionUser.KeyName(), &prop)
	require.NoError(t, err)
	t.Logf("Platform percentage change proposal submitted with ID %s in transaction %s", paramChangeTx.ProposalID, paramChangeTx.TxHash)

	require.Eventuallyf(t, func() bool {
		proposalInfo, err := xion.QueryProposal(ctx, paramChangeTx.ProposalID)
		if err != nil {
			require.NoError(t, err)
		} else {
			if proposalInfo.Status == cosmos.ProposalStatusVotingPeriod {
				return true
			}
			t.Logf("Waiting for proposal to enter voting status VOTING, current status: %s", proposalInfo.Status)
		}
		return false
	}, time.Second*11, time.Second, "failed to reach status VOTING after 11s")

	err = xion.VoteOnProposalAllValidators(ctx, paramChangeTx.ProposalID, cosmos.ProposalVoteYes)
	require.NoError(t, err)

	require.Eventuallyf(t, func() bool {
		proposalInfo, err := xion.QueryProposal(ctx, paramChangeTx.ProposalID)
		if err != nil {
			require.NoError(t, err)
		} else {
			if proposalInfo.Status == cosmos.ProposalStatusPassed {
				return true
			}
			t.Logf("Waiting for proposal to enter voting status PASSED, current status: %s", proposalInfo.Status)
		}
		return false
	}, time.Second*11, time.Second, "failed to reach status PASSED after 11s")

	// step 3: transfer and verify platform fees is extracted
	initialSendingBalance, err := xion.GetBalance(ctx, xionUser.FormattedAddress(), xion.Config().Denom)
	require.NoError(t, err)
	initialReceivingBalance, err := xion.GetBalance(ctx, recipientKeyAddress, xion.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, uint64(100), uint64(initialReceivingBalance))

	_, err = xion.FullNodes[0].ExecTx(ctx,
		xionUser.KeyName(),
		"xion", "send", xionUser.KeyName(),
		"--chain-id", xion.Config().ChainID,
		recipientKeyAddress, fmt.Sprintf("%d%s", 200, xion.Config().Denom),
	)
	require.NoError(t, err)

	postSendingBalance, err := xion.GetBalance(ctx, xionUser.FormattedAddress(), xion.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, uint64(initialSendingBalance-200), uint64(postSendingBalance))
	postReceivingBalance, err := xion.GetBalance(ctx, recipientKeyAddress, xion.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, uint64(290), uint64(postReceivingBalance))
}

func TestMintModuleInflationNoTransaction(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}

	t.Parallel()

	xion, ctx := BuildXionChain(t)
	// Query the mint module for the current inflation
	var inflation json.Number
	queryRes, _, err := xion.FullNodes[0].ExecQuery(ctx, "mint", "inflation")
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(queryRes, &inflation))
	inflationValue, err := inflation.Float64()
	t.Logf("Current inflation: %f", inflationValue)
	require.NoError(t, err, "inflation should be a float")
	// Make sure inflation is 0
	require.Equal(t, 0.0, inflationValue)

	// Query the mint module for inflation rate change
	var params = make(map[string]interface{})
	queryRes, _, err = xion.FullNodes[0].ExecQuery(ctx, "mint", "params")
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(queryRes, &params))
	inflationRateChange, err := dyno.GetString(params, "inflation_rate_change")
	require.NoError(t, err, "inflation_rate_change should be a string")
	inflationRateChangeValue, err := strconv.ParseFloat(inflationRateChange, 64)
	require.NoError(t, err, "inflation_rate_change should be convertible to float")
	t.Logf("Current inflation rate change: %f", inflationRateChangeValue)
	// Make sure inflation rate change is 0
	require.Equal(t, 0.0, inflationRateChangeValue)

	// Get the total bank supply
	jsonRes := make(map[string]interface{})
	queryRes, _, err = xion.FullNodes[0].ExecQuery(ctx, "bank", "total")
	require.NoError(t, err)

	require.NoError(t, json.Unmarshal(queryRes, &jsonRes))

	// Presuming we are the only denom on the chain
	totalSupply, err := dyno.GetSlice(jsonRes, "supply")
	require.NoError(t, err)
	xionCoin := totalSupply[0]
	require.NotEmpty(t, xionCoin)
	// Make sure we selected the uxion denom
	xionCoinDenom, err := dyno.GetString(xionCoin, "denom")
	require.NoError(t, err)
	require.Equal(t, xionCoinDenom, xion.Config().Denom)
	initialXionSupply, err := dyno.GetString(xionCoin, "amount")
	require.NoError(t, err)
	t.Logf("Initial Xion supply: %s", initialXionSupply)

	// Wait for some blocks and check if that supply stays the same
	chainHeight, _ := xion.Height(ctx)
	testutil.WaitForBlocks(ctx, int(chainHeight) + 10, xion)
	// Get the total bank supply
	currentResJson := make(map[string]interface{})
	currentSupplyRes, _, queryErr := xion.FullNodes[0].ExecQuery(ctx, "bank", "total")
	require.NoError(t, queryErr)
	require.NoError(t, json.Unmarshal(currentSupplyRes, &currentResJson))

	newTotalSupply, err := dyno.GetSlice(currentResJson, "supply")
	require.NoError(t, err)
	currentXionCoin := newTotalSupply[0]

	currentXionSupply, err := dyno.GetString(currentXionCoin, "amount")
	require.NoError(t, err)
	t.Logf("Current Xion supply: %s", currentXionSupply)

	require.Equal(t, initialXionSupply, currentXionSupply)
}