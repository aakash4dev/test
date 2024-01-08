package test

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"test/testutil/sample"
	testsimulation "test/x/test/simulation"
	"test/x/test/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = testsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateListof = "op_weight_msg_listof"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateListof int = 100

	opWeightMsgUpdateListof = "op_weight_msg_listof"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateListof int = 100

	opWeightMsgDeleteListof = "op_weight_msg_listof"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteListof int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	testGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ListofList: []types.Listof{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		ListofCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&testGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateListof int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateListof, &weightMsgCreateListof, nil,
		func(_ *rand.Rand) {
			weightMsgCreateListof = defaultWeightMsgCreateListof
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateListof,
		testsimulation.SimulateMsgCreateListof(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateListof int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateListof, &weightMsgUpdateListof, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateListof = defaultWeightMsgUpdateListof
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateListof,
		testsimulation.SimulateMsgUpdateListof(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteListof int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteListof, &weightMsgDeleteListof, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteListof = defaultWeightMsgDeleteListof
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteListof,
		testsimulation.SimulateMsgDeleteListof(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateListof,
			defaultWeightMsgCreateListof,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				testsimulation.SimulateMsgCreateListof(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateListof,
			defaultWeightMsgUpdateListof,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				testsimulation.SimulateMsgUpdateListof(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteListof,
			defaultWeightMsgDeleteListof,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				testsimulation.SimulateMsgDeleteListof(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
