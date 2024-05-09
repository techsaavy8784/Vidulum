package verify_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "vidulum/testutil/keeper"
	"vidulum/testutil/nullify"
	"vidulum/x/verify"
	"vidulum/x/verify/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ExternalAddressList: []types.ExternalAddress{
			{
				Vidulum: "0",
			},
			{
				Vidulum: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VerifyKeeper(t)
	verify.InitGenesis(ctx, *k, genesisState)
	got := verify.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ExternalAddressList, got.ExternalAddressList)
	// this line is used by starport scaffolding # genesis/test/assert
}
