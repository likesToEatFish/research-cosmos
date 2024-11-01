package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/vaults/types"
)

type msgServer struct {
	Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) UpdateParams(goCtx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if k.GetAuthority() != req.Authority {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.GetAuthority(), req.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.SetParams(ctx, req.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}

// Add new Active Collateral via gov
func (k msgServer) ActiveCollateral(ctx context.Context, msg *types.MsgActiveCollateral) (*types.MsgActiveCollateralResponse, error) {
	if k.authority != msg.Authority {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	err := k.ActiveCollateralAsset(ctx, msg.Denom, msg.MinCollateralRatio, msg.LiquidationRatio, msg.MaxDebt, msg.StabilityFee, msg.MintingFee, msg.LiquidationPenalty, int64(msg.OraclScript))
	if err != nil {
		return nil, err
	}

	return &types.MsgActiveCollateralResponse{}, nil
}

// Updates Collateral via gov
func (k msgServer) UpdatesCollateral(ctx context.Context, msg *types.MsgUpdatesCollateral) (*types.MsgUpdatesCollateralResponse, error) {
	if k.authority != msg.Authority {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	err := k.UpdatesCollateralAsset(ctx, msg.Denom, msg.MinCollateralRatio, msg.LiquidationRatio, msg.MaxDebt, msg.StabilityFee, msg.MintingFee, msg.LiquidationPenalty)
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdatesCollateralResponse{}, nil
}

// Create new vault, send Collateral and receive back an amount Minted
func (k msgServer) CreateVault(ctx context.Context, msg *types.MsgCreateVault) (*types.MsgCreateVaultResponse, error) {
	if !msg.Minted.IsPositive() {
		return nil, fmt.Errorf("invalid coin")
	}
	err := k.CreateNewVault(ctx, sdk.MustAccAddressFromBech32(msg.Owner), msg.Collateral, msg.Minted)
	if err != nil {
		return nil, err
	}
	return &types.MsgCreateVaultResponse{}, nil
}

// Send additional Collateral
func (k msgServer) Deposit(ctx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	if !msg.Amount.IsPositive() {
		return nil, fmt.Errorf("invalid coin")
	}

	err := k.DepositToVault(ctx, msg.VaultId, sdk.MustAccAddressFromBech32(msg.Sender), msg.Amount)
	if err != nil {
		return nil, err
	}
	return &types.MsgDepositResponse{}, nil
}

// Withdraw a amount Collateral, make sure the remaining Collateral value is still more than the loan amount
func (k msgServer) Withdraw(ctx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	if !msg.Amount.IsPositive() {
		return nil, fmt.Errorf("invalid coin")
	}

	err := k.WithdrawFromVault(ctx, msg.VaultId, sdk.MustAccAddressFromBech32(msg.Sender), msg.Amount)
	if err != nil {
		return nil, err
	}
	return &types.MsgWithdrawResponse{}, nil
}

// additional loan, collateral is still guaranteed
func (k msgServer) Mint(ctx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	if !msg.Amount.IsPositive() {
		return nil, fmt.Errorf("invalid coin")
	}

	err := k.MintCoin(ctx, msg.VaultId, sdk.MustAccAddressFromBech32(msg.Sender), msg.Amount)
	if err != nil {
		return nil, err
	}
	return &types.MsgMintResponse{}, nil
}

// repay part or all of a loan
func (k msgServer) Repay(ctx context.Context, msg *types.MsgRepay) (*types.MsgRepayResponse, error) {
	if !msg.Amount.IsPositive() {
		return nil, fmt.Errorf("invalid coin")
	}
	
	err := k.RepayDebt(ctx, msg.VaultId, sdk.MustAccAddressFromBech32(msg.Sender), msg.Amount)
	if err != nil {
		return nil, err
	}
	return &types.MsgRepayResponse{}, nil
}

// claim back the CollateralLocked, ensuring the debt is paid off
func (k msgServer) Close(ctx context.Context, msg *types.MsgClose) (*types.MsgCloseResponse, error) {
	vault, err := k.GetVault(ctx, msg.VaultId)
	if err != nil {
		return nil, fmt.Errorf("vault %d was not found", msg.VaultId)
	}
	
	err = k.CloseVault(ctx, msg.Sender, vault)
	if err != nil {
		return nil, err
	}
	return &types.MsgCloseResponse{}, nil
}
