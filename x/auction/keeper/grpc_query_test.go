package keeper_test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/require"

// 	"github.com/likesToEatFish/research-cosmos/x/oracle/types"
// 	keepertest "github.com/onomyprotocol/reserve/testutil/keeper"
// )

// func TestParamsQuery(t *testing.T) {
// 	keeper, ctx := keepertest.OracleKeeper(t)
// 	params := types.DefaultParams()
// 	require.NoError(t, keeper.SetParams(ctx, params))

// 	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
// 	require.NoError(t, err)
// 	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
// }
