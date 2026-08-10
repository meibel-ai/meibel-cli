package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/meibel-ai/meibel-cli/internal/tui"
)

var executionsGetRealtimeProgressCmd = &cobra.Command{
	Use:   "get-realtime-progress <execution-id>",
	Short: "Get Batch Realtime Progress",
	Long:  `Get Batch Realtime Progress

Arguments:
  execution-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel batches executions get-realtime-progress <execution-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		executionId := args[0]

		sp := tui.StartSpinner("Get Batch Realtime Progress")
		result, err := client.Batches.Executions.GetRealtimeProgress(ctx, executionId)
		sp.Stop()
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	executionsCmd.AddCommand(executionsGetRealtimeProgressCmd)

}
