package nameservice

import (
	"fmt"
	"github.com/arthaszeng/nameservice/x/nameservice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	WhoisRecords []types.Whois `json:"whois_records"`
}

func NewGenesisState(whoIsRecords []types.Whois) GenesisState {
	return GenesisState{WhoisRecords: whoIsRecords}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		WhoisRecords: []types.Whois{},
	}
}

func ValidateGenesis(genesisState GenesisState) error {
	for _, record := range genesisState.WhoisRecords {
		if record.Owner == nil {
			return fmt.Errorf("invalid WhoisRecord: Value: %s. Error: Missing Owner", record.Owner)
		}
		if record.Value == "" {
			return fmt.Errorf("invalid WhoisRecord: Owner: %s. Error: Missing Value", record.Value)
		}
		if record.Price == nil {
			return fmt.Errorf("invalid WhoisRecord: Value: %s. Error: Missing Price", record.Price)
		}
	}

	return nil
}

func InitGenesis(ctx sdk.Context, keeper Keeper, genesisState GenesisState) []abci.ValidatorUpdate {
	for _, record := range genesisState.WhoisRecords {
		keeper.SetWhois(ctx, record.Value, record)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	var records []types.Whois
	iterator := keeper.GetNamesIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {

		name := string(iterator.Key())
		whois := keeper.GetWhois(ctx, name)
		records = append(records, whois)

	}

	return GenesisState{WhoisRecords: records}
}
