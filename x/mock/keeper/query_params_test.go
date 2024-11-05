package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/likesToEatFish/research-cosmos/testutil/keeper"
	"github.com/likesToEatFish/research-cosmos/x/mock/keeper"
	"github.com/likesToEatFish/research-cosmos/x/mock/types"
)

func TestParamsQuery(t *testing.T) {
	k, ctx, _ := keepertest.MockKeeper(t)

	qs := keeper.NewQueryServerImpl(k)
	params := types.DefaultParams()
	require.NoError(t, k.Params.Set(ctx, params))

	response, err := qs.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
