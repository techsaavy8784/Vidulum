package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"vidulum/x/verify/types"
)

func CmdListExternalAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-external-address",
		Short: "list all externalAddress",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllExternalAddressRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ExternalAddressAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowExternalAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-external-address [vidulum]",
		Short: "shows a externalAddress",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argVidulum := args[0]

			params := &types.QueryGetExternalAddressRequest{
				Vidulum: argVidulum,
			}

			res, err := queryClient.ExternalAddress(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
