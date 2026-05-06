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
	instancesFailBlueprintInstanceData string
	instancesFailBlueprintInstanceInteractive bool
)

var instancesFailBlueprintInstanceCmd = &cobra.Command{
	Use:   "fail-blueprint <blueprint-instance-id>",
	Short: "Fail a Blueprint Instance",
	Long:  `This endpoint is used to mark a Blueprint Instance as failed. It will update the status of the Blueprint Instance to 'FAILED' and log the failure event.

Arguments:
  blueprint-instance-id: Unique identifier for the workflow instance`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints instances fail-blueprint <blueprint-instance-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		var body sdk.FailBlueprintInstanceRequest

		if instancesFailBlueprintInstanceData != "" {
			if err := json.Unmarshal([]byte(instancesFailBlueprintInstanceData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if instancesFailBlueprintInstanceInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Blueprints.Instances.FailBlueprintInstance(ctx, blueprintInstanceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	instancesCmd.AddCommand(instancesFailBlueprintInstanceCmd)

	instancesFailBlueprintInstanceCmd.Flags().StringVarP(&instancesFailBlueprintInstanceData, "data", "d", "", "JSON data for the request body")
	instancesFailBlueprintInstanceCmd.Flags().BoolVarP(&instancesFailBlueprintInstanceInteractive, "interactive", "i", false, "use interactive form input")
}
