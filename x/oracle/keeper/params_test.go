package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "reserve/testutil/keeper"
	"reserve/x/oracle/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.OracleKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
