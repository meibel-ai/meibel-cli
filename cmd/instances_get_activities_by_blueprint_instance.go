package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	instancesGetActivitiesByBlueprintInstanceOffset int64
	instancesGetActivitiesByBlueprintInstanceLimit int64
	instancesGetActivitiesByBlueprintInstanceSortBy string
	instancesGetActivitiesByBlueprintInstanceSortOrder string
)

var instancesGetActivitiesByBlueprintInstanceCmd = &cobra.Command{
	Use:   "get-activities-by-blueprint <blueprint-instance-id>",
	Short: "Get Activities By Blueprint Instance",
	Long:  `Get Activities By Blueprint Instance

Arguments:
  blueprint-instance-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints instances get-activities-by-blueprint <blueprint-instance-id> --offset=<value> --limit=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		opts := &sdk.GetActivitiesByBlueprintInstanceOptions{}
		if instancesGetActivitiesByBlueprintInstanceOffset != 0 {
			opts.Offset = &instancesGetActivitiesByBlueprintInstanceOffset
		}
		if instancesGetActivitiesByBlueprintInstanceLimit != 0 {
			opts.Limit = &instancesGetActivitiesByBlueprintInstanceLimit
		}
		if instancesGetActivitiesByBlueprintInstanceSortBy != "" {
			opts.SortBy = &instancesGetActivitiesByBlueprintInstanceSortBy
		}
		if instancesGetActivitiesByBlueprintInstanceSortOrder != "" {
			opts.SortOrder = &instancesGetActivitiesByBlueprintInstanceSortOrder
		}

		result, err := client.Blueprints.Instances.GetActivitiesByBlueprintInstance(ctx, blueprintInstanceId, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	instancesCmd.AddCommand(instancesGetActivitiesByBlueprintInstanceCmd)

	instancesGetActivitiesByBlueprintInstanceCmd.Flags().Int64VarP(&instancesGetActivitiesByBlueprintInstanceOffset, "offset", "", 0, "Number of items to skip")
	instancesGetActivitiesByBlueprintInstanceCmd.Flags().Int64VarP(&instancesGetActivitiesByBlueprintInstanceLimit, "limit", "", 10, "Maximum number of items to return")
	instancesGetActivitiesByBlueprintInstanceCmd.Flags().StringVarP(&instancesGetActivitiesByBlueprintInstanceSortBy, "sort-by", "", "", "Field to sort by")
	instancesGetActivitiesByBlueprintInstanceCmd.Flags().StringVarP(&instancesGetActivitiesByBlueprintInstanceSortOrder, "sort-order", "", "", "Sort order (asc or desc)")
}
