package cli

import (
	"fmt"
	"github.com/arthaszeng/nameservice/x/nameservice/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
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

func GetCmdQueryNames(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "names",
		Short: "names",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliContext := context.NewCLIContext()
			res, _, err := cliContext.QueryWithData(fmt.Sprintf("custom/%s/names", queryRoute), nil)

			if err != nil {
				fmt.Printf("Cannot get query of names \n")
				return nil
			}

			var output types.QueryResNames
			cdc.MustUnmarshalJSON(res, &output)
			return cliContext.PrintOutput(output)
		},
	}
}

func GetCmdQueryWhois(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "whois [name]",
		Short: "whois name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliContext := context.NewCLIContext().WithCodec(cdc)
			name := args[0]

			res, _, err := cliContext.QueryWithData(fmt.Sprintf("custom/%s/whois/%s", queryRoute, name), nil)

			if err != nil {
				fmt.Printf("Cannot whois name: %s\n", name)
				return nil
			}

			var output types.Whois
			cdc.MustUnmarshalJSON(res, &output)
			return cliContext.PrintOutput(output)
		},
	}
}

func GetCmdQueryResolve(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "resolve [name]",
		Short: "resolve name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliContext := context.NewCLIContext().WithCodec(cdc)
			name := args[0]

			res, _, err := cliContext.QueryWithData(fmt.Sprintf("custom/%s/resolve/%s", queryRoute, name), nil)

			if err != nil {
				fmt.Printf("Could not resolve name: %s\n", name)
				return nil
			}

			var output types.QueryResResolve
			cdc.MustUnmarshalJSON(res, &output)
			return cliContext.PrintOutput(output)
		},
	}
}
