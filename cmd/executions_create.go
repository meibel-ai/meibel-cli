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
	executionsCreateData string
	executionsCreateInteractive bool
)

var executionsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Batch Execution",
	Long:  `Create Batch Execution`,
	Example: "meibel batches executions create",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		var body sdk.CreateBatchExecutionRequest

		if executionsCreateData != "" {
			if err := json.Unmarshal([]byte(executionsCreateData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if executionsCreateInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Batches.Executions.Create(ctx, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	executionsCmd.AddCommand(executionsCreateCmd)

	executionsCreateCmd.Flags().StringVarP(&executionsCreateData, "data", "d", "", "JSON data for the request body")
	executionsCreateCmd.Flags().BoolVarP(&executionsCreateInteractive, "interactive", "i", false, "use interactive form input")
}
