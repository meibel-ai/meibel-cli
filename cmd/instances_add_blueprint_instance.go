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
	instancesAddBlueprintInstanceData string
	instancesAddBlueprintInstanceInteractive bool
)

var instancesAddBlueprintInstanceCmd = &cobra.Command{
	Use:   "add-blueprint",
	Short: "Add Blueprint Instance",
	Long:  `Add Blueprint Instance`,
	Example: "meibel blueprints instances add-blueprint",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		var body sdk.AddBlueprintInstanceRequest

		if instancesAddBlueprintInstanceData != "" {
			if err := json.Unmarshal([]byte(instancesAddBlueprintInstanceData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if instancesAddBlueprintInstanceInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Blueprints.Instances.AddBlueprintInstance(ctx, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	instancesCmd.AddCommand(instancesAddBlueprintInstanceCmd)

	instancesAddBlueprintInstanceCmd.Flags().StringVarP(&instancesAddBlueprintInstanceData, "data", "d", "", "JSON data for the request body")
	instancesAddBlueprintInstanceCmd.Flags().BoolVarP(&instancesAddBlueprintInstanceInteractive, "interactive", "i", false, "use interactive form input")
}
