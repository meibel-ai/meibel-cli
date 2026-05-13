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
	batchesCreateData string
	batchesCreateInteractive bool
)

var batchesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Batch Definition",
	Long:  `Create Batch Definition`,
	Example: "meibel batches create",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		var body sdk.CreateBatchDefinitionRequest

		if batchesCreateData != "" {
			if err := json.Unmarshal([]byte(batchesCreateData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if batchesCreateInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("Name").Description("Kebab-case label (non-unique within tenant)").Value(&body.Name),
					huh.NewInput().Title("AgentId").Description("AgentDefinition ID; resolved + pinned at creation time").Value(&body.AgentId),
					huh.NewInput().Title("InputDatasourceId").Description("Datasource holding the input Data Elements").Value(&body.InputDatasourceId),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.Batches.Create(ctx, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	batchesCmd.AddCommand(batchesCreateCmd)

	batchesCreateCmd.Flags().StringVarP(&batchesCreateData, "data", "d", "", "JSON data for the request body")
	batchesCreateCmd.Flags().BoolVarP(&batchesCreateInteractive, "interactive", "i", false, "use interactive form input")
}
