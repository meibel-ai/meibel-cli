package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var tagGetTagColumnInfoCmd = &cobra.Command{
	Use:   "get-column-info <datasource-id> <table-name> <column-name>",
	Short: "Get Tag Column Info",
	Long:  `Get Tag Column Info

Arguments:
  datasource-id: required
  table-name: required
  column-name: required`,
	Args:  cobra.ExactArgs(3),
	Example: "meibel datasources tag get-column-info <datasource-id> <table-name> <column-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		tableName := args[1]
		columnName := args[2]

		result, err := client.Datasources.Tag.GetTagColumnInfo(ctx, datasourceId, tableName, columnName)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tagCmd.AddCommand(tagGetTagColumnInfoCmd)

}
