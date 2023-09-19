package dongtramcam_test

import (
	"testing"

	keepertest "github.com/DongCoNY/research-cosmos/testutil/keeper"
	"github.com/DongCoNY/research-cosmos/testutil/nullify"
	"github.com/DongCoNY/research-cosmos/x/dongtramcam"
	"github.com/DongCoNY/research-cosmos/x/dongtramcam/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DongtramcamKeeper(t)
	dongtramcam.InitGenesis(ctx, *k, genesisState)
	got := dongtramcam.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
