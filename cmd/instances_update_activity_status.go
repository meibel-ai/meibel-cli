package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	instancesUpdateActivityStatusUpdatedStatusValue string
)

var instancesUpdateActivityStatusCmd = &cobra.Command{
	Use:   "update-activity-status <blueprint-instance-id> <activity-id>",
	Short: "Update Activity Status",
	Long:  `Update Activity Status

Arguments:
  blueprint-instance-id: required
  activity-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel blueprints instances update-activity-status <blueprint-instance-id> <activity-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]
		activityId := args[1]

		err := client.Blueprints.Instances.UpdateActivityStatus(ctx, blueprintInstanceId, activityId, sdk.ActivityStatus(instancesUpdateActivityStatusUpdatedStatusValue))
		if err != nil {
			return err
		}

		fmt.Println("Success")
		return nil
	},
}

func init() {
	instancesCmd.AddCommand(instancesUpdateActivityStatusCmd)

	instancesUpdateActivityStatusCmd.Flags().StringVarP(&instancesUpdateActivityStatusUpdatedStatusValue, "updated-status-value", "", "", "The updated-status-value parameter")
	instancesUpdateActivityStatusCmd.MarkFlagRequired("updated-status-value")
}
