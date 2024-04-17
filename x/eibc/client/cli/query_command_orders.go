package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/dymensionxyz/dymension/v3/x/eibc/types"
)

func CmdListDemandOrdersByStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-demand-orders [status]",
		Short: "List all demand orders with a specific status",
		Long:  `Query demand orders filtered by status. Examples of status include "pending", "finalized", and "reverted".`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			request := &types.QueryDemandOrdersRequest{
				Status: args[0],
			}

			rollappId, err := cmd.Flags().GetString("rollapp_id")
			if err == nil {
				request.RollappId = rollappId
			}

			orderType, err := cmd.Flags().GetString("type")
			if err == nil {
				request.Type = orderType
			}

			limit, err := cmd.Flags().GetInt32("limit")
			if err == nil {
				request.Limit = limit
			}

			res, err := queryClient.DemandOrders(context.Background(), request)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String("rollapp_id", "", "Optional rollapp_id")
	cmd.Flags().String("type", "", "Optional type")
	cmd.Flags().Int32("limit", 0, "Optional limit")

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
