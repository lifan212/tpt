package tpt_test

import (
	"testing"

	keepertest "tpt/testutil/keeper"
	"tpt/testutil/nullify"
	tpt "tpt/x/tpt/module"
	"tpt/x/tpt/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TptKeeper(t)
	tpt.InitGenesis(ctx, k, genesisState)
	got := tpt.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
