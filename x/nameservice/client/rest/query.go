package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"net/http"
)

func resolveWhoisHandler(cliContext context.CLIContext, storeName string) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		name := vars[restName]

		res, _, err := cliContext.QueryWithData(fmt.Sprintf("custom/%s/whois/%s", storeName, name), nil)
		if err != nil {
			rest.WriteErrorResponse(writer, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(writer, cliContext, res)
	}
}

func resolveNameHandler(cliContext context.CLIContext, storeName string) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		name := vars[restName]

		res, _, err := cliContext.QueryWithData(fmt.Sprintf("custom/%s/resolve/%s", storeName, name), nil)
		if err != nil {
			rest.WriteErrorResponse(writer, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(writer, cliContext, res)
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
