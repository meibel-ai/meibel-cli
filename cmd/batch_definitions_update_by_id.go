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
	batchDefinitionsUpdateByIdData string
	batchDefinitionsUpdateByIdInteractive bool
)

var batchDefinitionsUpdateByIdCmd = &cobra.Command{
	Use:   "update-by-id <definition-id>",
	Short: "Update Batch Definition By Id",
	Long:  `Update Batch Definition By Id

Arguments:
  definition-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel batch-definitions update-by-id <definition-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		definitionId := args[0]

		var body sdk.UpdateBatchDefinitionRequest

		if batchDefinitionsUpdateByIdData != "" {
			if err := json.Unmarshal([]byte(batchDefinitionsUpdateByIdData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if batchDefinitionsUpdateByIdInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.BatchDefinitions.UpdateById(ctx, definitionId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	batchDefinitionsCmd.AddCommand(batchDefinitionsUpdateByIdCmd)

	batchDefinitionsUpdateByIdCmd.Flags().StringVarP(&batchDefinitionsUpdateByIdData, "data", "d", "", "JSON data for the request body")
	batchDefinitionsUpdateByIdCmd.Flags().BoolVarP(&batchDefinitionsUpdateByIdInteractive, "interactive", "i", false, "use interactive form input")
}
