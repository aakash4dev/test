package test_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "test/testutil/keeper"
	"test/testutil/nullify"
	"test/x/test"
	"test/x/test/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ListofList: []types.Listof{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ListofCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TestKeeper(t)
	test.InitGenesis(ctx, *k, genesisState)
	got := test.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ListofList, got.ListofList)
	require.Equal(t, genesisState.ListofCount, got.ListofCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
