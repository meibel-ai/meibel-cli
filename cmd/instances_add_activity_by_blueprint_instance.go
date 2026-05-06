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
	instancesAddActivityByBlueprintInstanceData string
	instancesAddActivityByBlueprintInstanceInteractive bool
)

var instancesAddActivityByBlueprintInstanceCmd = &cobra.Command{
	Use:   "add-activity-by-blueprint <blueprint-instance-id>",
	Short: "Add Activity By Blueprint Instance",
	Long:  `Add Activity By Blueprint Instance

Arguments:
  blueprint-instance-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints instances add-activity-by-blueprint <blueprint-instance-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		var body sdk.AddActivityRequest

		if instancesAddActivityByBlueprintInstanceData != "" {
			if err := json.Unmarshal([]byte(instancesAddActivityByBlueprintInstanceData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if instancesAddActivityByBlueprintInstanceInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("ActivityType").Description("").Value(&body.ActivityType),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.Blueprints.Instances.AddActivityByBlueprintInstance(ctx, blueprintInstanceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	instancesCmd.AddCommand(instancesAddActivityByBlueprintInstanceCmd)

	instancesAddActivityByBlueprintInstanceCmd.Flags().StringVarP(&instancesAddActivityByBlueprintInstanceData, "data", "d", "", "JSON data for the request body")
	instancesAddActivityByBlueprintInstanceCmd.Flags().BoolVarP(&instancesAddActivityByBlueprintInstanceInteractive, "interactive", "i", false, "use interactive form input")
}
