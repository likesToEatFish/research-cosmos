package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/likesToEatFish/research-cosmos/x/oracle/types"
	keepertest "github.com/onomyprotocol/reserve/testutil/keeper"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.OracleKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
