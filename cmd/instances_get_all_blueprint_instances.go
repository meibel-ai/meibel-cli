package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	instancesGetAllBlueprintInstancesIncludeChildren bool
	instancesGetAllBlueprintInstancesIncludeActivities bool
	instancesGetAllBlueprintInstancesIncludeEvents bool
	instancesGetAllBlueprintInstancesOffset int64
	instancesGetAllBlueprintInstancesLimit int64
	instancesGetAllBlueprintInstancesSortBy string
	instancesGetAllBlueprintInstancesSortOrder string
)

var instancesGetAllBlueprintInstancesCmd = &cobra.Command{
	Use:   "get-all-blueprint",
	Short: "Get All Blueprint Instances",
	Long:  `Get All Blueprint Instances`,
	Example: "meibel blueprints instances get-all-blueprint --include-children=<value> --include-activities=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := &sdk.GetAllBlueprintInstancesOptions{}
		if instancesGetAllBlueprintInstancesIncludeChildren {
			opts.IncludeChildren = &instancesGetAllBlueprintInstancesIncludeChildren
		}
		if instancesGetAllBlueprintInstancesIncludeActivities {
			opts.IncludeActivities = &instancesGetAllBlueprintInstancesIncludeActivities
		}
		if instancesGetAllBlueprintInstancesIncludeEvents {
			opts.IncludeEvents = &instancesGetAllBlueprintInstancesIncludeEvents
		}
		if instancesGetAllBlueprintInstancesOffset != 0 {
			opts.Offset = &instancesGetAllBlueprintInstancesOffset
		}
		if instancesGetAllBlueprintInstancesLimit != 0 {
			opts.Limit = &instancesGetAllBlueprintInstancesLimit
		}
		if instancesGetAllBlueprintInstancesSortBy != "" {
			opts.SortBy = &instancesGetAllBlueprintInstancesSortBy
		}
		if instancesGetAllBlueprintInstancesSortOrder != "" {
			opts.SortOrder = &instancesGetAllBlueprintInstancesSortOrder
		}

		result, err := client.Blueprints.Instances.GetAllBlueprintInstances(ctx, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	instancesCmd.AddCommand(instancesGetAllBlueprintInstancesCmd)

	instancesGetAllBlueprintInstancesCmd.Flags().BoolVarP(&instancesGetAllBlueprintInstancesIncludeChildren, "include-children", "", false, "The include-children parameter")
	instancesGetAllBlueprintInstancesCmd.Flags().BoolVarP(&instancesGetAllBlueprintInstancesIncludeActivities, "include-activities", "", false, "The include-activities parameter")
	instancesGetAllBlueprintInstancesCmd.Flags().BoolVarP(&instancesGetAllBlueprintInstancesIncludeEvents, "include-events", "", false, "The include-events parameter")
	instancesGetAllBlueprintInstancesCmd.Flags().Int64VarP(&instancesGetAllBlueprintInstancesOffset, "offset", "", 0, "Number of items to skip")
	instancesGetAllBlueprintInstancesCmd.Flags().Int64VarP(&instancesGetAllBlueprintInstancesLimit, "limit", "", 10, "Maximum number of items to return")
	instancesGetAllBlueprintInstancesCmd.Flags().StringVarP(&instancesGetAllBlueprintInstancesSortBy, "sort-by", "", "", "Field to sort by")
	instancesGetAllBlueprintInstancesCmd.Flags().StringVarP(&instancesGetAllBlueprintInstancesSortOrder, "sort-order", "", "", "Sort order (asc or desc)")
}
