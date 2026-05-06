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
	blueprintsCreateBlueprintData string
	blueprintsCreateBlueprintInteractive bool
)

var blueprintsCreateBlueprintCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Blueprint",
	Long:  `Create Blueprint`,
	Example: "meibel blueprints create",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		var body sdk.AddBlueprintRequest

		if blueprintsCreateBlueprintData != "" {
			if err := json.Unmarshal([]byte(blueprintsCreateBlueprintData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if blueprintsCreateBlueprintInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("Name").Description("").Value(&body.Name),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.Blueprints.CreateBlueprint(ctx, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsCmd.AddCommand(blueprintsCreateBlueprintCmd)

	blueprintsCreateBlueprintCmd.Flags().StringVarP(&blueprintsCreateBlueprintData, "data", "d", "", "JSON data for the request body")
	blueprintsCreateBlueprintCmd.Flags().BoolVarP(&blueprintsCreateBlueprintInteractive, "interactive", "i", false, "use interactive form input")
}
