package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	dataelementsGetDataElementsOffset int64
	dataelementsGetDataElementsLimit int64
	dataelementsGetDataElementsSortBy string
	dataelementsGetDataElementsSortOrder string
)

var dataelementsGetDataElementsCmd = &cobra.Command{
	Use:   "get-data-elements <datasource-id>",
	Short: "Get Data Elements",
	Long:  `Get Data Elements

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources dataelements get-data-elements <datasource-id> --offset=<value> --limit=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		opts := &sdk.GetDataElementsOptions{}
		if dataelementsGetDataElementsOffset != 0 {
			opts.Offset = &dataelementsGetDataElementsOffset
		}
		if dataelementsGetDataElementsLimit != 0 {
			opts.Limit = &dataelementsGetDataElementsLimit
		}
		if dataelementsGetDataElementsSortBy != "" {
			opts.SortBy = &dataelementsGetDataElementsSortBy
		}
		if dataelementsGetDataElementsSortOrder != "" {
			opts.SortOrder = &dataelementsGetDataElementsSortOrder
		}

		result, err := client.Datasources.Dataelements.GetDataElements(ctx, datasourceId, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	dataelementsCmd.AddCommand(dataelementsGetDataElementsCmd)

	dataelementsGetDataElementsCmd.Flags().Int64VarP(&dataelementsGetDataElementsOffset, "offset", "", 0, "Number of items to skip")
	dataelementsGetDataElementsCmd.Flags().Int64VarP(&dataelementsGetDataElementsLimit, "limit", "", 10, "Maximum number of items to return")
	dataelementsGetDataElementsCmd.Flags().StringVarP(&dataelementsGetDataElementsSortBy, "sort-by", "", "", "Field to sort by")
	dataelementsGetDataElementsCmd.Flags().StringVarP(&dataelementsGetDataElementsSortOrder, "sort-order", "", "", "Sort order (asc or desc)")
}
