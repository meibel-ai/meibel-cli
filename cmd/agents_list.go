package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	agentsListOffset int64
	agentsListLimit string
)

var agentsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Agents",
	Long:  `List Agents`,
	Example: "meibel agents list --offset=<value> --limit=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := &sdk.AgentsListOptions{}
		if agentsListOffset != 0 {
			opts.Offset = &agentsListOffset
		}
		if agentsListLimit != "" {
			opts.Limit = &agentsListLimit
		}

		iter := client.Agents.List(ctx, opts)

		var items []interface{}
		for iter.Next(ctx) {
			items = append(items, iter.Item())
		}
		if err := iter.Err(); err != nil {
			return err
		}

		return output.Print(items)
	},
}

func init() {
	agentsCmd.AddCommand(agentsListCmd)

	agentsListCmd.Flags().Int64VarP(&agentsListOffset, "offset", "", 0, "Number of items to skip")
	agentsListCmd.Flags().StringVarP(&agentsListLimit, "limit", "", "", "Maximum number of items to return")
}
