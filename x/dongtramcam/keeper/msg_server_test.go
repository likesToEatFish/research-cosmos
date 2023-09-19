package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/DongCoNY/research-cosmos/testutil/keeper"
	"github.com/DongCoNY/research-cosmos/x/dongtramcam/keeper"
	"github.com/DongCoNY/research-cosmos/x/dongtramcam/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DongtramcamKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
