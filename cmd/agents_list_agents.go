package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	agentsListAgentsOffset int64
	agentsListAgentsLimit string
)

var agentsListAgentsCmd = &cobra.Command{
	Use:   "list",
	Short: "List Agents",
	Long:  `List Agents`,
	Example: "meibel agents list --offset=<value> --limit=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := &sdk.ListAgentsOptions{}
		if agentsListAgentsOffset != 0 {
			opts.Offset = &agentsListAgentsOffset
		}
		if agentsListAgentsLimit != "" {
			opts.Limit = &agentsListAgentsLimit
		}

		iter := client.Agents.ListAgents(ctx, opts)

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
	agentsCmd.AddCommand(agentsListAgentsCmd)

	agentsListAgentsCmd.Flags().Int64VarP(&agentsListAgentsOffset, "offset", "", 0, "Number of items to skip")
	agentsListAgentsCmd.Flags().StringVarP(&agentsListAgentsLimit, "limit", "", "", "Maximum number of items to return")
}
