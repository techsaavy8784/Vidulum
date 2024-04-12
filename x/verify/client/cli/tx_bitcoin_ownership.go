package cli

import (
	"strconv"

	"vidulum/x/verify/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBitcoinOwnership() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bitcoin-ownership [address] [signature] [message]",
		Short: "Broadcast message bitcoin-ownership",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]
			argSignature := args[1]
			argMessage := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBitcoinOwnership(
				clientCtx.GetFromAddress().String(),
				argAddress,
				argSignature,
				argMessage,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
