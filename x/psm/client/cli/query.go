package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/likesToEatFish/research-cosmos/x/psm/types"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd() *cobra.Command {
	// Group dao queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2, // nolint:gomnd
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCheckParams())
	cmd.AddCommand(CmdCheckStablecoin())
	cmd.AddCommand(CmdGetAllStablecoin())
	return cmd
}

func CmdCheckStablecoin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stablecoin [denom]",
		Short: "shows info price, fee in, fee out, swapable quantity of the stablecoin",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Stablecoin(context.Background(), &types.QueryStablecoinRequest{Denom: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetAllStablecoin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-stablecoin",
		Short: "shows info price, fee in, fee out, swapable quantity of all stablecoin",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AllStablecoin(context.Background(), &types.QueryAllStablecoinRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdCheckParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params [denom]",
		Short: "shows params module psm",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
