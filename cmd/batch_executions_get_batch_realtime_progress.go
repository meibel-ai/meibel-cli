package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var batchExecutionsGetBatchRealtimeProgressCmd = &cobra.Command{
	Use:   "get-realtime-progress <execution-id>",
	Short: "Get Batch Realtime Progress",
	Long:  `Get Batch Realtime Progress

Arguments:
  execution-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel batch-executions get-realtime-progress <execution-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		executionId := args[0]

		result, err := client.BatchExecutions.GetBatchRealtimeProgress(ctx, executionId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	batchExecutionsCmd.AddCommand(batchExecutionsGetBatchRealtimeProgressCmd)

}
