package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	agentsSessionsListOffset int64
	agentsSessionsListLimit int64
	agentsSessionsListSortBy string
	agentsSessionsListSortOrder string
	agentsSessionsListStatus string
)

var agentsSessionsListCmd = &cobra.Command{
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

		opts := &sdk.AgentsSessionsListOptions{}
		if agentsSessionsListOffset != 0 {
			opts.Offset = &agentsSessionsListOffset
		}
		if agentsSessionsListLimit != 0 {
			opts.Limit = &agentsSessionsListLimit
		}
		if agentsSessionsListSortBy != "" {
			opts.SortBy = &agentsSessionsListSortBy
		}
		if agentsSessionsListSortOrder != "" {
			opts.SortOrder = &agentsSessionsListSortOrder
		}
		if agentsSessionsListStatus != "" {
			opts.Status = &agentsSessionsListStatus
		}

		iter := client.Agents.Sessions.List(ctx, agentId, opts)

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
	agentsSessionsCmd.AddCommand(agentsSessionsListCmd)

	agentsSessionsListCmd.Flags().Int64VarP(&agentsSessionsListOffset, "offset", "", 0, "Number of items to skip")
	agentsSessionsListCmd.Flags().Int64VarP(&agentsSessionsListLimit, "limit", "", 10, "Maximum number of items to return")
	agentsSessionsListCmd.Flags().StringVarP(&agentsSessionsListSortBy, "sort-by", "", "start_time", "Field to sort by: start_time, status")
	agentsSessionsListCmd.Flags().StringVarP(&agentsSessionsListSortOrder, "sort-order", "", "desc", "Sort order: asc or desc")
	agentsSessionsListCmd.Flags().StringVarP(&agentsSessionsListStatus, "status", "", "", "Filter by execution status: RUNNING, COMPLETED, FAILED, CANCELED, TERMINATED")
}
