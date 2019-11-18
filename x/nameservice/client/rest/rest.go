package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

func RegisterRoutes(cliContext context.CLIContext, router mux.Router, storeName string) {
	router.HandleFunc(fmt.Sprintf("%s/names", storeName), namesHandler(cliContext, storeName)).Methods("GET")
	router.HandleFunc(fmt.Sprintf("%s/names", storeName), buyNameHander(cliContext, storeName)).Methods("POST")
	router.HandleFunc(fmt.Sprintf("%s/names", storeName), setNameHander(cliContext, storeName)).Methods("PUT")
	router.HandleFunc(fmt.Sprintf("%s/names/name", storeName), resolveNameHandler(cliContext, storeName)).Methods("GET")
	router.HandleFunc(fmt.Sprintf("%s/names/whois", storeName), resolveWhoisHandler(cliContext, storeName)).Methods("GET")
}
