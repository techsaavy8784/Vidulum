package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "vidulum/testutil/keeper"
	"vidulum/x/verify/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.VerifyKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
