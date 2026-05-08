package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var tablesListColumnsCmd = &cobra.Command{
	Use:   "list-columns <datasource-id> <table-name>",
	Short: "List Columns",
	Long:  `List Columns

Arguments:
  datasource-id: required
  table-name: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources tables list-columns <datasource-id> <table-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		tableName := args[1]

		result, err := client.Datasources.Tables.ListColumns(ctx, datasourceId, tableName)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tablesCmd.AddCommand(tablesListColumnsCmd)

}
