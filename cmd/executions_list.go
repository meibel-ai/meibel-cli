package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	executionsListInputDatasourceId string
	executionsListOffset int64
	executionsListLimit string
	executionsListSortBy string
	executionsListSortOrder string
)

var executionsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Batch Executions",
	Long:  `List Batch Executions`,
	Example: "meibel batches executions list --input-datasource-id=<value> --offset=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := &sdk.ExecutionsListOptions{}
		if executionsListInputDatasourceId != "" {
			opts.InputDatasourceId = &executionsListInputDatasourceId
		}
		if executionsListOffset != 0 {
			opts.Offset = &executionsListOffset
		}
		if executionsListLimit != "" {
			opts.Limit = &executionsListLimit
		}
		if executionsListSortBy != "" {
			opts.SortBy = &executionsListSortBy
		}
		if executionsListSortOrder != "" {
			opts.SortOrder = &executionsListSortOrder
		}

		result, err := client.Batches.Executions.List(ctx, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	executionsCmd.AddCommand(executionsListCmd)

	executionsListCmd.Flags().StringVarP(&executionsListInputDatasourceId, "input-datasource-id", "", "", "Filter by input datasource ID")
	executionsListCmd.Flags().Int64VarP(&executionsListOffset, "offset", "", 0, "The offset parameter")
	executionsListCmd.Flags().StringVarP(&executionsListLimit, "limit", "", "", "The limit parameter")
	executionsListCmd.Flags().StringVarP(&executionsListSortBy, "sort-by", "", "start_time", "Field to sort by: start_time, status")
	executionsListCmd.Flags().StringVarP(&executionsListSortOrder, "sort-order", "", "desc", "The sort-order parameter")
}
