package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var executionsRetryFailedItemsCmd = &cobra.Command{
	Use:   "retry-failed-items <execution-id>",
	Short: "Retry Failed Items",
	Long:  `Retry Failed Items

Arguments:
  execution-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel batches executions retry-failed-items <execution-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		executionId := args[0]

		result, err := client.Batches.Executions.RetryFailedItems(ctx, executionId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	executionsCmd.AddCommand(executionsRetryFailedItemsCmd)

}
