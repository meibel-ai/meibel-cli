package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
	"golang.org/x/term"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	batchExecutionsUpdateByIdData string
	batchExecutionsUpdateByIdInteractive bool
)

var batchExecutionsUpdateByIdCmd = &cobra.Command{
	Use:   "update-by-id <execution-id>",
	Short: "Update Batch Execution By Id",
	Long:  `Update Batch Execution By Id

Arguments:
  execution-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel batch-executions update-by-id <execution-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		executionId := args[0]

		var body sdk.UpdateBatchExecutionRequest

		if batchExecutionsUpdateByIdData != "" {
			if err := json.Unmarshal([]byte(batchExecutionsUpdateByIdData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if batchExecutionsUpdateByIdInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("AdditionalProperties").Description(""),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.BatchExecutions.UpdateById(ctx, executionId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	batchExecutionsCmd.AddCommand(batchExecutionsUpdateByIdCmd)

	batchExecutionsUpdateByIdCmd.Flags().StringVarP(&batchExecutionsUpdateByIdData, "data", "d", "", "JSON data for the request body")
	batchExecutionsUpdateByIdCmd.Flags().BoolVarP(&batchExecutionsUpdateByIdInteractive, "interactive", "i", false, "use interactive form input")
}
