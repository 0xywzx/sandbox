package cosmos_test

import (
	"testing"

	keepertest "github.com/0xywzx/cosmos/testutil/keeper"
	"github.com/0xywzx/cosmos/testutil/nullify"
	"github.com/0xywzx/cosmos/x/cosmos"
	"github.com/0xywzx/cosmos/x/cosmos/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmosKeeper(t)
	cosmos.InitGenesis(ctx, *k, genesisState)
	got := cosmos.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
