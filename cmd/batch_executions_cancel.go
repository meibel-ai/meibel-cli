package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var batchExecutionsCancelCmd = &cobra.Command{
	Use:   "cancel <execution-id>",
	Short: "Cancel Batch Execution",
	Long:  `Cancel Batch Execution

Arguments:
  execution-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel batch-executions cancel <execution-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		executionId := args[0]

		result, err := client.BatchExecutions.Cancel(ctx, executionId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	batchExecutionsCmd.AddCommand(batchExecutionsCancelCmd)

}
