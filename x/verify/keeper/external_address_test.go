package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "vidulum/testutil/keeper"
	"vidulum/testutil/nullify"
	"vidulum/x/verify/keeper"
	"vidulum/x/verify/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNExternalAddress(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ExternalAddress {
	items := make([]types.ExternalAddress, n)
	for i := range items {
		items[i].Vidulum = strconv.Itoa(i)

		keeper.SetExternalAddress(ctx, items[i])
	}
	return items
}

func TestExternalAddressGet(t *testing.T) {
	keeper, ctx := keepertest.VerifyKeeper(t)
	items := createNExternalAddress(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetExternalAddress(ctx,
			item.Vidulum,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestExternalAddressRemove(t *testing.T) {
	keeper, ctx := keepertest.VerifyKeeper(t)
	items := createNExternalAddress(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveExternalAddress(ctx,
			item.Vidulum,
		)
		_, found := keeper.GetExternalAddress(ctx,
			item.Vidulum,
		)
		require.False(t, found)
	}
}

func TestExternalAddressGetAll(t *testing.T) {
	keeper, ctx := keepertest.VerifyKeeper(t)
	items := createNExternalAddress(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllExternalAddress(ctx)),
	)
}
