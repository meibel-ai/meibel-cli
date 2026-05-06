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
	blueprintsUpdateBlueprintData string
	blueprintsUpdateBlueprintInteractive bool
)

var blueprintsUpdateBlueprintCmd = &cobra.Command{
	Use:   "update <blueprint-id>",
	Short: "Update Blueprint",
	Long:  `Update Blueprint

Arguments:
  blueprint-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints update <blueprint-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintId := args[0]

		var body sdk.UpdateBlueprintRequest

		if blueprintsUpdateBlueprintData != "" {
			if err := json.Unmarshal([]byte(blueprintsUpdateBlueprintData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if blueprintsUpdateBlueprintInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Blueprints.UpdateBlueprint(ctx, blueprintId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsCmd.AddCommand(blueprintsUpdateBlueprintCmd)

	blueprintsUpdateBlueprintCmd.Flags().StringVarP(&blueprintsUpdateBlueprintData, "data", "d", "", "JSON data for the request body")
	blueprintsUpdateBlueprintCmd.Flags().BoolVarP(&blueprintsUpdateBlueprintInteractive, "interactive", "i", false, "use interactive form input")
}
