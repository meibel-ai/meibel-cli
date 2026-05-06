package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	instancesUpdateBlueprintInstanceStatusUpdatedStatusValue string
	instancesUpdateBlueprintInstanceStatusWorkflowRunId string
)

var instancesUpdateBlueprintInstanceStatusCmd = &cobra.Command{
	Use:   "update-blueprint-status <blueprint-instance-id>",
	Short: "Update Blueprint Instance Status",
	Long:  `Update Blueprint Instance Status

Arguments:
  blueprint-instance-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints instances update-blueprint-status <blueprint-instance-id> --workflow-run-id=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		opts := &sdk.UpdateBlueprintInstanceStatusOptions{}
		if instancesUpdateBlueprintInstanceStatusWorkflowRunId != "" {
			opts.WorkflowRunId = &instancesUpdateBlueprintInstanceStatusWorkflowRunId
		}

		err := client.Blueprints.Instances.UpdateBlueprintInstanceStatus(ctx, blueprintInstanceId, sdk.BlueprintInstanceStatus(instancesUpdateBlueprintInstanceStatusUpdatedStatusValue), opts)
		if err != nil {
			return err
		}

		fmt.Println("Success")
		return nil
	},
}

func init() {
	instancesCmd.AddCommand(instancesUpdateBlueprintInstanceStatusCmd)

	instancesUpdateBlueprintInstanceStatusCmd.Flags().StringVarP(&instancesUpdateBlueprintInstanceStatusUpdatedStatusValue, "updated-status-value", "", "", "The updated-status-value parameter")
	instancesUpdateBlueprintInstanceStatusCmd.MarkFlagRequired("updated-status-value")
	instancesUpdateBlueprintInstanceStatusCmd.Flags().StringVarP(&instancesUpdateBlueprintInstanceStatusWorkflowRunId, "workflow-run-id", "", "", "The workflow-run-id parameter")
}
