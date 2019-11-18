package cli

import (
	"github.com/arthaszeng/nameservice/x/nameservice/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        "nameservice",
		Short:                      "Name service transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(client.PostCommands(
		GetCmdSetName(cdc),
		GetCmdBuyName(cdc),
	)...)

	return txCmd
}

func GetCmdBuyName(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "buy-name [name] [amount]",
		Short: "bid for existing name or claim a new name",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliContext := context.NewCLIContext().WithCodec(cdc)
			txBuilder := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			name := args[0]
			bid := args[1]
			accAddresses := cliContext.GetFromAddress()

			coins, err := sdk.ParseCoins(bid)
			if err != nil {
				return err
			}

			msgBuyName := types.NewMsgBuyName(name, coins, accAddresses)
			err = msgBuyName.ValidateBasic()

			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliContext, txBuilder, []sdk.Msg{msgBuyName})
		},
	}
}

func GetCmdSetName(cdc *codec.Codec) *cobra.Command {

}
