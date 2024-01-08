package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "test/testutil/keeper"
	"test/testutil/nullify"
	"test/x/test/keeper"
	"test/x/test/types"
)

func createNListof(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Listof {
	items := make([]types.Listof, n)
	for i := range items {
		items[i].Id = keeper.AppendListof(ctx, items[i])
	}
	return items
}

func TestListofGet(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNListof(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetListof(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestListofRemove(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNListof(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveListof(ctx, item.Id)
		_, found := keeper.GetListof(ctx, item.Id)
		require.False(t, found)
	}
}

func TestListofGetAll(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNListof(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllListof(ctx)),
	)
}

func TestListofCount(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	items := createNListof(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetListofCount(ctx))
}
