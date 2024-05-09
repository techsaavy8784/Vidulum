package verify

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"vidulum/testutil/sample"
	verifysimulation "vidulum/x/verify/simulation"
	"vidulum/x/verify/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = verifysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgBitcoinOwnership = "op_weight_msg_bitcoin_ownership"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBitcoinOwnership int = 100

	opWeightMsgSolanaOwnership = "op_weight_msg_solana_ownership"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSolanaOwnership int = 100

	opWeightMsgEthereumOwnership = "op_weight_msg_ethereum_ownership"
	// TODO: Determine the simulation weight value
	defaultWeightMsgEthereumOwnership int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	verifyGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&verifyGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgBitcoinOwnership int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBitcoinOwnership, &weightMsgBitcoinOwnership, nil,
		func(_ *rand.Rand) {
			weightMsgBitcoinOwnership = defaultWeightMsgBitcoinOwnership
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBitcoinOwnership,
		verifysimulation.SimulateMsgBitcoinOwnership(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSolanaOwnership int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSolanaOwnership, &weightMsgSolanaOwnership, nil,
		func(_ *rand.Rand) {
			weightMsgSolanaOwnership = defaultWeightMsgSolanaOwnership
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSolanaOwnership,
		verifysimulation.SimulateMsgSolanaOwnership(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgEthereumOwnership int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgEthereumOwnership, &weightMsgEthereumOwnership, nil,
		func(_ *rand.Rand) {
			weightMsgEthereumOwnership = defaultWeightMsgEthereumOwnership
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEthereumOwnership,
		verifysimulation.SimulateMsgEthereumOwnership(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
