package rest

import (
	"fmt"
	"github.com/arthaszeng/nameservice/x/nameservice/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes(cliContext context.CLIContext, router mux.Router, storeName string) {
	router.HandleFunc(fmt.Sprintf("%s/names", storeName), namesHandler(cliContext, storeName)).Methods("GET")
	router.HandleFunc(fmt.Sprintf("%s/names", storeName), buyNameHandler(cliContext)).Methods("POST")
	router.HandleFunc(fmt.Sprintf("%s/names", storeName), setNameHandler(cliContext)).Methods("PUT")
	router.HandleFunc(fmt.Sprintf("%s/names/name", storeName), resolveNameHandler(cliContext, storeName)).Methods("GET")
	router.HandleFunc(fmt.Sprintf("%s/names/whois", storeName), resolveWhoisHandler(cliContext, storeName)).Methods("GET")
}

type setNameReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Name    string       `json:"name"`
	Value   string       `json:"value"`
	Owner   string       `json:"owner"`
}

func setNameHandler(cliContext context.CLIContext) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var setNameRequest setNameReq
		if !rest.ReadRESTReq(writer, request, cliContext.Codec, &setNameRequest) {
			rest.WriteErrorResponse(writer, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := setNameRequest.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(writer) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(setNameRequest.Owner)
		if err != nil {
			rest.WriteErrorResponse(writer, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgSetName(setNameRequest.Name, setNameRequest.Value, addr)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(writer, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(writer, cliContext, baseReq, []sdk.Msg{msg})
	}
}

type buyNameReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Name    string       `json:"name"`
	Amount  string       `json:"amount"`
	Buyer   string       `json:"buyer"`
}

func buyNameHandler(cliContext context.CLIContext) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var buyNameRequest buyNameReq

		if !rest.ReadRESTReq(writer, request, cliContext.Codec, &buyNameRequest) {
			rest.WriteErrorResponse(writer, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := buyNameRequest.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(writer) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(buyNameRequest.Buyer)
		if err != nil {
			rest.WriteErrorResponse(writer, http.StatusBadRequest, err.Error())
			return
		}

		coins, err := sdk.ParseCoins(buyNameRequest.Amount)
		if err != nil {
			rest.WriteErrorResponse(writer, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgBuyName(buyNameRequest.Name, coins, addr)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(writer, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(writer, cliContext, baseReq, []sdk.Msg{msg})
	}
}

func namesHandler(cliContext context.CLIContext, storeName string) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		res, _, err := cliContext.QueryWithData(fmt.Sprintf("custom/%s/namse", storeName), nil)
		if err != nil {
			rest.WriteErrorResponse(writer, http.StatusNotFound, err.Error())
		}

		rest.PostProcessResponse(writer, cliContext, res)
	}
}
