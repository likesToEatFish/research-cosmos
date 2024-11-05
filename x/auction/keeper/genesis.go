package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/likesToEatFish/research-cosmos/x/auction/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
	k.LastestAuctionPeriod = ctx.BlockTime().Unix()
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	return genesis
}
