package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var instancesGetActivityByBlueprintInstanceCmd = &cobra.Command{
	Use:   "get-activity-by-blueprint <blueprint-instance-id> <activity-id>",
	Short: "Get Activity By Blueprint Instance",
	Long:  `Get Activity By Blueprint Instance

Arguments:
  blueprint-instance-id: required
  activity-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel blueprints instances get-activity-by-blueprint <blueprint-instance-id> <activity-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]
		activityId := args[1]

		result, err := client.Blueprints.Instances.GetActivityByBlueprintInstance(ctx, blueprintInstanceId, activityId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	instancesCmd.AddCommand(instancesGetActivityByBlueprintInstanceCmd)

}
