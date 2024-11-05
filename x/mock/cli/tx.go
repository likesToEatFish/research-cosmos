package cli

import (
	"context"
	"cosmossdk.io/math"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/likesToEatFish/research-cosmos/x/mock/types"
)

// GetTxCmd returns the transaction commands for this module.
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewSetPriceCmd(),
	)

	return cmd
}

func NewSetPriceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-price [denom] [price]",
		Args:  cobra.ExactArgs(2),
		Short: "set price for denom ",
		Long: `set price for denomn.

			Example:
			$ onomyd tx mockoracel set-price usdt 1 
	`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			price := math.LegacyMustNewDecFromStr(args[1])
			addr := clientCtx.GetFromAddress()
			msg := types.NewSetPrice(args[0], addr.String(), price)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the oracle module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdGetPrice(),
	)
	return cmd
}

func CmdGetPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-price [denom1] [denom2]",
		Short: "shows info price denom",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			msg := types.NewQueryPrice(args[0], args[1])
			res, err := queryClient.Price(context.Background(), &msg)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
