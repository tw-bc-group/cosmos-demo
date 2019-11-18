package nameservice

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/abci/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

type AppModuleBasic struct {
}

func (appModuleBasic AppModuleBasic) Name() string {
	panic("implement me")
}

func (appModuleBasic AppModuleBasic) RegisterCodec(*codec.Codec) {
	panic("implement me")
}

func (appModuleBasic AppModuleBasic) DefaultGenesis() json.RawMessage {
	panic("implement me")
}

func (appModuleBasic AppModuleBasic) ValidateGenesis(json.RawMessage) error {
	panic("implement me")
}

func (appModuleBasic AppModuleBasic) RegisterRESTRoutes(context.CLIContext, *mux.Router) {
	panic("implement me")
}

func (appModuleBasic AppModuleBasic) GetTxCmd(*codec.Codec) *cobra.Command {
	panic("implement me")
}

func (appModuleBasic AppModuleBasic) GetQueryCmd(*codec.Codec) *cobra.Command {
	panic("implement me")
}

type AppModule struct {
	AppModuleBasic
	keeper     Keeper
	coinKeeper bank.Keeper
}

func (appModule AppModule) Name() string {
	panic("implement me")
}

func (appModule AppModule) RegisterCodec(*codec.Codec) {
	panic("implement me")
}

func (appModule AppModule) DefaultGenesis() json.RawMessage {
	panic("implement me")
}

func (appModule AppModule) ValidateGenesis(json.RawMessage) error {
	panic("implement me")
}

func (appModule AppModule) RegisterRESTRoutes(context.CLIContext, *mux.Router) {
	panic("implement me")
}

func (appModule AppModule) GetTxCmd(*codec.Codec) *cobra.Command {
	panic("implement me")
}

func (appModule AppModule) GetQueryCmd(*codec.Codec) *cobra.Command {
	panic("implement me")
}

func (appModule AppModule) InitGenesis(sdk.Context, json.RawMessage) []types.ValidatorUpdate {
	panic("implement me")
}

func (appModule AppModule) ExportGenesis(sdk.Context) json.RawMessage {
	panic("implement me")
}

func (appModule AppModule) RegisterInvariants(sdk.InvariantRegistry) {
	panic("implement me")
}

func (appModule AppModule) Route() string {
	panic("implement me")
}

func (appModule AppModule) NewHandler() sdk.Handler {
	panic("implement me")
}

func (appModule AppModule) QuerierRoute() string {
	panic("implement me")
}

func (appModule AppModule) NewQuerierHandler() sdk.Querier {
	panic("implement me")
}

func (appModule AppModule) BeginBlock(sdk.Context, types.RequestBeginBlock) {
	panic("implement me")
}

func (appModule AppModule) EndBlock(sdk.Context, types.RequestEndBlock) []types.ValidatorUpdate {
	panic("implement me")
}
