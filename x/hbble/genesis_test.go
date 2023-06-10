package hbble_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/walload/hbble/testutil/keeper"
	"github.com/walload/hbble/testutil/nullify"
	"github.com/walload/hbble/x/hbble"
	"github.com/walload/hbble/x/hbble/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HbbleKeeper(t)
	hbble.InitGenesis(ctx, *k, genesisState)
	got := hbble.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
