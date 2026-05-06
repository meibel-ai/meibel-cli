package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	tagGetAllTagColumnInfoOffset int64
	tagGetAllTagColumnInfoLimit int64
	tagGetAllTagColumnInfoSortBy string
	tagGetAllTagColumnInfoSortOrder string
)

var tagGetAllTagColumnInfoCmd = &cobra.Command{
	Use:   "get-all-column-info <datasource-id> <table-name>",
	Short: "Get All Tag Column Info",
	Long:  `Get All Tag Column Info

Arguments:
  datasource-id: required
  table-name: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources tag get-all-column-info <datasource-id> <table-name> --offset=<value> --limit=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		tableName := args[1]

		opts := &sdk.GetAllTagColumnInfoOptions{}
		if tagGetAllTagColumnInfoOffset != 0 {
			opts.Offset = &tagGetAllTagColumnInfoOffset
		}
		if tagGetAllTagColumnInfoLimit != 0 {
			opts.Limit = &tagGetAllTagColumnInfoLimit
		}
		if tagGetAllTagColumnInfoSortBy != "" {
			opts.SortBy = &tagGetAllTagColumnInfoSortBy
		}
		if tagGetAllTagColumnInfoSortOrder != "" {
			opts.SortOrder = &tagGetAllTagColumnInfoSortOrder
		}

		result, err := client.Datasources.Tag.GetAllTagColumnInfo(ctx, datasourceId, tableName, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tagCmd.AddCommand(tagGetAllTagColumnInfoCmd)

	tagGetAllTagColumnInfoCmd.Flags().Int64VarP(&tagGetAllTagColumnInfoOffset, "offset", "", 0, "Number of items to skip")
	tagGetAllTagColumnInfoCmd.Flags().Int64VarP(&tagGetAllTagColumnInfoLimit, "limit", "", 10, "Maximum number of items to return")
	tagGetAllTagColumnInfoCmd.Flags().StringVarP(&tagGetAllTagColumnInfoSortBy, "sort-by", "", "", "Field to sort by")
	tagGetAllTagColumnInfoCmd.Flags().StringVarP(&tagGetAllTagColumnInfoSortOrder, "sort-order", "", "", "Sort order (asc or desc)")
}
