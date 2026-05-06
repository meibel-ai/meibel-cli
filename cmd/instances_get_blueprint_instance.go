package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	instancesGetBlueprintInstanceIncludeChildren bool
	instancesGetBlueprintInstanceIncludeActivities bool
	instancesGetBlueprintInstanceIncludeEvents bool
)

var instancesGetBlueprintInstanceCmd = &cobra.Command{
	Use:   "get-blueprint <blueprint-instance-id>",
	Short: "Get Blueprint Instance",
	Long:  `Get Blueprint Instance

Arguments:
  blueprint-instance-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints instances get-blueprint <blueprint-instance-id> --include-children=<value> --include-activities=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		opts := &sdk.GetBlueprintInstanceOptions{}
		if instancesGetBlueprintInstanceIncludeChildren {
			opts.IncludeChildren = &instancesGetBlueprintInstanceIncludeChildren
		}
		if instancesGetBlueprintInstanceIncludeActivities {
			opts.IncludeActivities = &instancesGetBlueprintInstanceIncludeActivities
		}
		if instancesGetBlueprintInstanceIncludeEvents {
			opts.IncludeEvents = &instancesGetBlueprintInstanceIncludeEvents
		}

		result, err := client.Blueprints.Instances.GetBlueprintInstance(ctx, blueprintInstanceId, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	instancesCmd.AddCommand(instancesGetBlueprintInstanceCmd)

	instancesGetBlueprintInstanceCmd.Flags().BoolVarP(&instancesGetBlueprintInstanceIncludeChildren, "include-children", "", false, "The include-children parameter")
	instancesGetBlueprintInstanceCmd.Flags().BoolVarP(&instancesGetBlueprintInstanceIncludeActivities, "include-activities", "", false, "The include-activities parameter")
	instancesGetBlueprintInstanceCmd.Flags().BoolVarP(&instancesGetBlueprintInstanceIncludeEvents, "include-events", "", false, "The include-events parameter")
}
