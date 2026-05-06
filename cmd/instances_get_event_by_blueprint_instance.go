package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var instancesGetEventByBlueprintInstanceCmd = &cobra.Command{
	Use:   "get-event-by-blueprint <blueprint-instance-id> <event-id>",
	Short: "Get Event By Blueprint Instance",
	Long:  `Get Event By Blueprint Instance

Arguments:
  blueprint-instance-id: required
  event-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel blueprints instances get-event-by-blueprint <blueprint-instance-id> <event-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]
		eventId := args[1]

		result, err := client.Blueprints.Instances.GetEventByBlueprintInstance(ctx, blueprintInstanceId, eventId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	instancesCmd.AddCommand(instancesGetEventByBlueprintInstanceCmd)

}
