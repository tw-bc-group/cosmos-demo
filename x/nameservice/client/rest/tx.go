package rest

import (
	"github.com/arthaszeng/nameservice/x/nameservice/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"net/http"
)

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
