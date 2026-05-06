package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var (
	instancesCompleteBlueprintInstanceData string
	instancesCompleteBlueprintInstanceInteractive bool
)

var instancesCompleteBlueprintInstanceCmd = &cobra.Command{
	Use:   "complete-blueprint <blueprint-instance-id>",
	Short: "Complete a Blueprint Instance",
	Long:  `This endpoint is used to mark a Blueprint Instance as completed. It will update the status of the Blueprint Instance to 'COMPLETED' and log the completion event.

Arguments:
  blueprint-instance-id: Unique identifier for the workflow instance`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints instances complete-blueprint <blueprint-instance-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		var body interface{}

		if instancesCompleteBlueprintInstanceData != "" {
			if err := json.Unmarshal([]byte(instancesCompleteBlueprintInstanceData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else {
			return fmt.Errorf("--data flag required (interactive form not available for this type)")
		}

		result, err := client.Blueprints.Instances.CompleteBlueprintInstance(ctx, blueprintInstanceId, &body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	instancesCmd.AddCommand(instancesCompleteBlueprintInstanceCmd)

	instancesCompleteBlueprintInstanceCmd.Flags().StringVarP(&instancesCompleteBlueprintInstanceData, "data", "d", "", "JSON data for the request body")
	instancesCompleteBlueprintInstanceCmd.Flags().BoolVarP(&instancesCompleteBlueprintInstanceInteractive, "interactive", "i", false, "use interactive form input")
}
