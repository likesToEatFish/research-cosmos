package keeper_test

import (
	"testing"

	testkeeper "github.com/DongCoNY/research-cosmos/testutil/keeper"
	"github.com/DongCoNY/research-cosmos/x/dongtramcam/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DongtramcamKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
