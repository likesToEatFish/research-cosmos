package mock_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/test/mock/testutil/keeper"
	"github.com/test/mock/testutil/nullify"
	mock "github.com/test/mock/x/mock/module"
	"github.com/test/mock/x/mock/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx, _ := keepertest.MockKeeper(t)
	err := mock.InitGenesis(ctx, k, genesisState)
	require.NoError(t, err)
	got, err := mock.ExportGenesis(ctx, k)
	require.NoError(t, err)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.Params, got.Params)
	// this line is used by starport scaffolding # genesis/test/assert
}
