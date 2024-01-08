package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ListofList: []Listof{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in listof
	listofIdMap := make(map[uint64]bool)
	listofCount := gs.GetListofCount()
	for _, elem := range gs.ListofList {
		if _, ok := listofIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for listof")
		}
		if elem.Id >= listofCount {
			return fmt.Errorf("listof id should be lower or equal than the last id")
		}
		listofIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
