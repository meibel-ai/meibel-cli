package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	agentsSessionsListSessionsOffset int64
	agentsSessionsListSessionsLimit string
	agentsSessionsListSessionsSortBy string
	agentsSessionsListSessionsSortOrder string
	agentsSessionsListSessionsStatus string
)

var agentsSessionsListSessionsCmd = &cobra.Command{
	Use:   "list <agent-id>",
	Short: "List Sessions",
	Long:  `List Sessions

Arguments:
  agent-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel agents agents-sessions list <agent-id> --offset=<value> --limit=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		agentId := args[0]

		opts := &sdk.ListSessionsOptions{}
		if agentsSessionsListSessionsOffset != 0 {
			opts.Offset = &agentsSessionsListSessionsOffset
		}
		if agentsSessionsListSessionsLimit != "" {
			opts.Limit = &agentsSessionsListSessionsLimit
		}
		if agentsSessionsListSessionsSortBy != "" {
			opts.SortBy = &agentsSessionsListSessionsSortBy
		}
		if agentsSessionsListSessionsSortOrder != "" {
			opts.SortOrder = &agentsSessionsListSessionsSortOrder
		}
		if agentsSessionsListSessionsStatus != "" {
			opts.Status = &agentsSessionsListSessionsStatus
		}

		iter := client.Agents.Sessions.ListSessions(ctx, agentId, opts)

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
	agentsSessionsCmd.AddCommand(agentsSessionsListSessionsCmd)

	agentsSessionsListSessionsCmd.Flags().Int64VarP(&agentsSessionsListSessionsOffset, "offset", "", 0, "Number of items to skip")
	agentsSessionsListSessionsCmd.Flags().StringVarP(&agentsSessionsListSessionsLimit, "limit", "", "", "Maximum number of items to return")
	agentsSessionsListSessionsCmd.Flags().StringVarP(&agentsSessionsListSessionsSortBy, "sort-by", "", "start_time", "Field to sort by: start_time, status")
	agentsSessionsListSessionsCmd.Flags().StringVarP(&agentsSessionsListSessionsSortOrder, "sort-order", "", "desc", "Sort order: asc or desc")
	agentsSessionsListSessionsCmd.Flags().StringVarP(&agentsSessionsListSessionsStatus, "status", "", "", "Filter by execution status: RUNNING, COMPLETED, FAILED, CANCELED, TERMINATED")
}
