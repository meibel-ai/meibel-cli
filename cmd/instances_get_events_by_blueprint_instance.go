package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	instancesGetEventsByBlueprintInstanceOffset int64
	instancesGetEventsByBlueprintInstanceLimit int64
	instancesGetEventsByBlueprintInstanceSortBy string
	instancesGetEventsByBlueprintInstanceSortOrder string
)

var instancesGetEventsByBlueprintInstanceCmd = &cobra.Command{
	Use:   "get-events-by-blueprint <blueprint-instance-id>",
	Short: "Get Events By Blueprint Instance",
	Long:  `Get Events By Blueprint Instance

Arguments:
  blueprint-instance-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints instances get-events-by-blueprint <blueprint-instance-id> --offset=<value> --limit=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		opts := &sdk.GetEventsByBlueprintInstanceOptions{}
		if instancesGetEventsByBlueprintInstanceOffset != 0 {
			opts.Offset = &instancesGetEventsByBlueprintInstanceOffset
		}
		if instancesGetEventsByBlueprintInstanceLimit != 0 {
			opts.Limit = &instancesGetEventsByBlueprintInstanceLimit
		}
		if instancesGetEventsByBlueprintInstanceSortBy != "" {
			opts.SortBy = &instancesGetEventsByBlueprintInstanceSortBy
		}
		if instancesGetEventsByBlueprintInstanceSortOrder != "" {
			opts.SortOrder = &instancesGetEventsByBlueprintInstanceSortOrder
		}

		result, err := client.Blueprints.Instances.GetEventsByBlueprintInstance(ctx, blueprintInstanceId, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	instancesCmd.AddCommand(instancesGetEventsByBlueprintInstanceCmd)

	instancesGetEventsByBlueprintInstanceCmd.Flags().Int64VarP(&instancesGetEventsByBlueprintInstanceOffset, "offset", "", 0, "Number of items to skip")
	instancesGetEventsByBlueprintInstanceCmd.Flags().Int64VarP(&instancesGetEventsByBlueprintInstanceLimit, "limit", "", 10, "Maximum number of items to return")
	instancesGetEventsByBlueprintInstanceCmd.Flags().StringVarP(&instancesGetEventsByBlueprintInstanceSortBy, "sort-by", "", "", "Field to sort by")
	instancesGetEventsByBlueprintInstanceCmd.Flags().StringVarP(&instancesGetEventsByBlueprintInstanceSortOrder, "sort-order", "", "", "Sort order (asc or desc)")
}
