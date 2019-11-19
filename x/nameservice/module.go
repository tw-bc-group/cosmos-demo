package nameservice

import (
	"encoding/json"
	"github.com/arthaszeng/nameservice/x/nameservice/client/cli"
	"github.com/arthaszeng/nameservice/x/nameservice/client/rest"
	"github.com/arthaszeng/nameservice/x/nameservice/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

type AppModuleBasic struct {
}

func (appModuleBasic AppModuleBasic) Name() string {
	return types.ModuleName
}

func (appModuleBasic AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	types.RegisterCodec(cdc)
}

func (appModuleBasic AppModuleBasic) DefaultGenesis() json.RawMessage {
	return types.ModuleCdc.MustMarshalJSON(DefaultGenesisState())
}

func (appModuleBasic AppModuleBasic) ValidateGenesis(data json.RawMessage) error {
	var genesisState GenesisState
	err := types.ModuleCdc.UnmarshalJSON(data, &genesisState)

	if err != nil {
		return err
	}

	return ValidateGenesis(genesisState)
}

func (appModuleBasic AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, router *mux.Router) {
	rest.RegisterRoutes(ctx, router, types.ModuleName)
}

func (appModuleBasic AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetTxCmd(types.ModuleName, cdc)
}

func (appModuleBasic AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetQueryCmd(types.ModuleName, cdc)
}

type AppModule struct {
	AppModuleBasic
	keeper     Keeper
	coinKeeper bank.Keeper
}

func NewAppModule(k Keeper, bankKeeper bank.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         k,
		coinKeeper:     bankKeeper,
	}
}

func (AppModule) Name() string {
	return types.ModuleName
}

func (appModule AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState GenesisState
	types.ModuleCdc.MustUnmarshalJSON(data, &genesisState)
	return InitGenesis(ctx, appModule.keeper, genesisState)
}

func (appModule AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	gs := ExportGenesis(ctx, appModule.keeper)
	return types.ModuleCdc.MustMarshalJSON(gs)
}

func (appModule AppModule) RegisterInvariants(sdk.InvariantRegistry) {
	//
}

func (appModule AppModule) Route() string {
	return types.ModuleName
}

func (appModule AppModule) NewHandler() sdk.Handler {
	return NewHandler(appModule.keeper)
}

func (appModule AppModule) QuerierRoute() string {
	return types.ModuleName
}

func (appModule AppModule) NewQuerierHandler() sdk.Querier {
	return NewQuerier(appModule.keeper)
}

func (appModule AppModule) BeginBlock(sdk.Context, abci.RequestBeginBlock) {
	//
}

func (appModule AppModule) EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
