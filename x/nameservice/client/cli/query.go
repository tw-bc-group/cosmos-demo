package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        "nameservice",
		Short:                      "Query command for the name service",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	queryCmd.AddCommand(client.GetCommands(
		GetCmdQueryResolve(storeKey, cdc),
		GetCmdQueryWhois(storeKey, cdc),
		GetCmdQueryNames(storeKey, cdc),
	)...)

	return queryCmd
}
