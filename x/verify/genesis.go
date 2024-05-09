package verify

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"vidulum/x/verify/keeper"
	"vidulum/x/verify/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the externalAddress
	for _, elem := range genState.ExternalAddressList {
		k.SetExternalAddress(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ExternalAddressList = k.GetAllExternalAddress(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
