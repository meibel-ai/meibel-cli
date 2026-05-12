package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var batchExecutionsGetByIdCmd = &cobra.Command{
	Use:   "get-by-id <execution-id>",
	Short: "Get Batch Execution By Id",
	Long:  `Get Batch Execution By Id

Arguments:
  execution-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel batch-executions get-by-id <execution-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		executionId := args[0]

		result, err := client.BatchExecutions.GetById(ctx, executionId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	batchExecutionsCmd.AddCommand(batchExecutionsGetByIdCmd)

}
