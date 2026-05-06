package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	tagGetAllTagTableInfoOffset int64
	tagGetAllTagTableInfoLimit int64
	tagGetAllTagTableInfoSortBy string
	tagGetAllTagTableInfoSortOrder string
)

var tagGetAllTagTableInfoCmd = &cobra.Command{
	Use:   "get-all-table-info <datasource-id>",
	Short: "Get All Tag Table Info",
	Long:  `Get All Tag Table Info

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources tag get-all-table-info <datasource-id> --offset=<value> --limit=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		opts := &sdk.GetAllTagTableInfoOptions{}
		if tagGetAllTagTableInfoOffset != 0 {
			opts.Offset = &tagGetAllTagTableInfoOffset
		}
		if tagGetAllTagTableInfoLimit != 0 {
			opts.Limit = &tagGetAllTagTableInfoLimit
		}
		if tagGetAllTagTableInfoSortBy != "" {
			opts.SortBy = &tagGetAllTagTableInfoSortBy
		}
		if tagGetAllTagTableInfoSortOrder != "" {
			opts.SortOrder = &tagGetAllTagTableInfoSortOrder
		}

		result, err := client.Datasources.Tag.GetAllTagTableInfo(ctx, datasourceId, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tagCmd.AddCommand(tagGetAllTagTableInfoCmd)

	tagGetAllTagTableInfoCmd.Flags().Int64VarP(&tagGetAllTagTableInfoOffset, "offset", "", 0, "Number of items to skip")
	tagGetAllTagTableInfoCmd.Flags().Int64VarP(&tagGetAllTagTableInfoLimit, "limit", "", 10, "Maximum number of items to return")
	tagGetAllTagTableInfoCmd.Flags().StringVarP(&tagGetAllTagTableInfoSortBy, "sort-by", "", "", "Field to sort by")
	tagGetAllTagTableInfoCmd.Flags().StringVarP(&tagGetAllTagTableInfoSortOrder, "sort-order", "", "", "Sort order (asc or desc)")
}
