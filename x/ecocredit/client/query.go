package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// QueryCmd returns the parent command for all x/data CLI query commands
func QueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Args:                       cobra.ExactArgs(1),
		Use:                        ecocredit.ModuleName,
		Short:                      "Querying commands for the ecocredit module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	cmd.AddCommand(
		queryClassInfo(),
		queryBatchInfo(),
		queryBalance(),
		querySupply(),
		queryPrecision(),
	)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func print(ctx client.Context, res proto.Message, err error) error {
	if err != nil {
		return err
	}
	return ctx.PrintOutput(res)
}

func createClient(cmd *cobra.Command) (ecocredit.QueryClient, client.Context, error) {
	ctx := client.GetClientContextFromCmd(cmd)
	ctx, err := client.ReadQueryCommandFlags(ctx, cmd.Flags())
	return ecocredit.NewQueryClient(ctx), ctx, err
}

func queryClassInfo() *cobra.Command {
	return &cobra.Command{
		Use:   "class_info [class_id]",
		Short: "Select a class info",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := createClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.ClassInfo(cmd.Context(), &ecocredit.QueryClassInfoRequest{
				ClassId: args[0],
			})
			return print(ctx, res, err)
		},
	}
}

func queryBatchInfo() *cobra.Command {
	return &cobra.Command{
		Use:   "batch_info [batch_denom]",
		Short: "Select the credit issuance batch info",
		Long:  "Select the credit issuance batch info based on the bach_denom (ID)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := createClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.BatchInfo(cmd.Context(), &ecocredit.QueryBatchInfoRequest{
				BatchDenom: args[0],
			})
			return print(ctx, res, err)
		},
	}
}

func queryBalance() *cobra.Command {
	return &cobra.Command{
		Use:   "balance [batch_denom] [account]",
		Short: "Select the tradable and retired balance of the credit batch",
		Long:  "Select the credit tradable and retired balance of the credit batch denom (ID) and account address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := createClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.Balance(cmd.Context(), &ecocredit.QueryBalanceRequest{
				BatchDenom: args[0], Account: args[1],
			})
			return print(ctx, res, err)
		},
	}
}

func querySupply() *cobra.Command {
	return &cobra.Command{
		Use:   "supply [batch_denom]",
		Short: "Select the tradable and retired supply of the credit batch",
		Long:  "Select the tradable and retired supply of the credit batch batch denom (ID)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := createClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.Supply(cmd.Context(), &ecocredit.QuerySupplyRequest{
				BatchDenom: args[0],
			})
			return print(ctx, res, err)
		},
	}
}

func queryPrecision() *cobra.Command {
	return &cobra.Command{
		Use:   "precision [batch_denom]",
		Short: "Select the maximum length of the fractional part of credits in the given batch",
		Long:  "Select the maximum length of the fractional part of credits in the given batch. The precision tells what is the minimum unit of a credit.\nExample: a decimal number 12.345 has fractional part length equal 3. A precision=5 means that the minimum unit we can trade is 0.00001",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := createClient(cmd)
			if err != nil {
				return err
			}
			res, err := c.Precision(cmd.Context(), &ecocredit.QueryPrecisionRequest{
				BatchDenom: args[0],
			})
			return print(ctx, res, err)
		},
	}
}
