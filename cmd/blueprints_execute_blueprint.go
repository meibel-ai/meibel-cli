package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
	"golang.org/x/term"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	blueprintsExecuteBlueprintData string
	blueprintsExecuteBlueprintInteractive bool
)

var blueprintsExecuteBlueprintCmd = &cobra.Command{
	Use:   "execute <blueprint-id>",
	Short: "Execute Blueprint",
	Long:  `Execute Blueprint

Arguments:
  blueprint-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints execute <blueprint-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintId := args[0]

		var body sdk.ExecuteBlueprintRequest

		if blueprintsExecuteBlueprintData != "" {
			if err := json.Unmarshal([]byte(blueprintsExecuteBlueprintData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if blueprintsExecuteBlueprintInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Blueprints.ExecuteBlueprint(ctx, blueprintId, &body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsCmd.AddCommand(blueprintsExecuteBlueprintCmd)

	blueprintsExecuteBlueprintCmd.Flags().StringVarP(&blueprintsExecuteBlueprintData, "data", "d", "", "JSON data for the request body")
	blueprintsExecuteBlueprintCmd.Flags().BoolVarP(&blueprintsExecuteBlueprintInteractive, "interactive", "i", false, "use interactive form input")
}
