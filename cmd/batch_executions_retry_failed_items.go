package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var batchExecutionsRetryFailedItemsCmd = &cobra.Command{
	Use:   "retry-failed-items <execution-id>",
	Short: "Retry Failed Items",
	Long:  `Retry Failed Items

Arguments:
  execution-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel batch-executions retry-failed-items <execution-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		executionId := args[0]

		result, err := client.BatchExecutions.RetryFailedItems(ctx, executionId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	batchExecutionsCmd.AddCommand(batchExecutionsRetryFailedItemsCmd)

}
