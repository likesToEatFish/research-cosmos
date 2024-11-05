package mock_test

import (
	"testing"

	keepertest "github.com/likesToEatFish/research-cosmos/testutil/keeper"
	"github.com/likesToEatFish/research-cosmos/testutil/nullify"
	mock "github.com/likesToEatFish/research-cosmos/x/mock/module"
	"github.com/likesToEatFish/research-cosmos/x/mock/types"
	"github.com/stretchr/testify/require"
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
