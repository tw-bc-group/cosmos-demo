package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

var restName = "name"

func RegisterRoutes(cliContext context.CLIContext, router *mux.Router, storeName string) {
	router.HandleFunc(fmt.Sprintf("%s/names", storeName), namesHandler(cliContext, storeName)).Methods("GET")
	router.HandleFunc(fmt.Sprintf("%s/names", storeName), buyNameHandler(cliContext)).Methods("POST")
	router.HandleFunc(fmt.Sprintf("%s/names", storeName), setNameHandler(cliContext)).Methods("PUT")
	router.HandleFunc(fmt.Sprintf("%s/names/{%s}", storeName, restName), resolveNameHandler(cliContext, storeName)).Methods("GET")
	router.HandleFunc(fmt.Sprintf("%s/names/{%s}/whois", storeName, restName), resolveWhoisHandler(cliContext, storeName)).Methods("GET")
}
