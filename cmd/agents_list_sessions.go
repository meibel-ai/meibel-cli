package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	agentsListSessionsOffset int64
	agentsListSessionsLimit string
	agentsListSessionsSortBy string
	agentsListSessionsSortOrder string
	agentsListSessionsStatus string
)

var agentsListSessionsCmd = &cobra.Command{
	Use:   "list-sessions <agent-id>",
	Short: "List Sessions",
	Long:  `List Sessions

Arguments:
  agent-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel agents list-sessions <agent-id> --offset=<value> --limit=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		agentId := args[0]

		opts := &sdk.ListSessionsOptions{}
		if agentsListSessionsOffset != 0 {
			opts.Offset = &agentsListSessionsOffset
		}
		if agentsListSessionsLimit != "" {
			opts.Limit = &agentsListSessionsLimit
		}
		if agentsListSessionsSortBy != "" {
			opts.SortBy = &agentsListSessionsSortBy
		}
		if agentsListSessionsSortOrder != "" {
			opts.SortOrder = &agentsListSessionsSortOrder
		}
		if agentsListSessionsStatus != "" {
			opts.Status = &agentsListSessionsStatus
		}

		iter := client.Agents.ListSessions(ctx, agentId, opts)

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
	agentsCmd.AddCommand(agentsListSessionsCmd)

	agentsListSessionsCmd.Flags().Int64VarP(&agentsListSessionsOffset, "offset", "", 0, "Number of items to skip")
	agentsListSessionsCmd.Flags().StringVarP(&agentsListSessionsLimit, "limit", "", "", "Maximum number of items to return")
	agentsListSessionsCmd.Flags().StringVarP(&agentsListSessionsSortBy, "sort-by", "", "start_time", "Field to sort by: start_time, status")
	agentsListSessionsCmd.Flags().StringVarP(&agentsListSessionsSortOrder, "sort-order", "", "desc", "Sort order: asc or desc")
	agentsListSessionsCmd.Flags().StringVarP(&agentsListSessionsStatus, "status", "", "", "Filter by execution status: RUNNING, COMPLETED, FAILED, CANCELED, TERMINATED")
}
