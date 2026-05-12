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
	batchExecutionsCreateData string
	batchExecutionsCreateInteractive bool
)

var batchExecutionsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Batch Execution",
	Long:  `Create Batch Execution`,
	Example: "meibel batch-executions create",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		var body sdk.CreateBatchExecutionRequest

		if batchExecutionsCreateData != "" {
			if err := json.Unmarshal([]byte(batchExecutionsCreateData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if batchExecutionsCreateInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.BatchExecutions.Create(ctx, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	batchExecutionsCmd.AddCommand(batchExecutionsCreateCmd)

	batchExecutionsCreateCmd.Flags().StringVarP(&batchExecutionsCreateData, "data", "d", "", "JSON data for the request body")
	batchExecutionsCreateCmd.Flags().BoolVarP(&batchExecutionsCreateInteractive, "interactive", "i", false, "use interactive form input")
}
