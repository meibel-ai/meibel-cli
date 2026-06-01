package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var tablesListColumnsCmd = &cobra.Command{
	Use:   "list-columns <table-name> <datasource-id>",
	Short: "List Columns",
	Long:  `List Columns

Arguments:
  table-name: required
  datasource-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources tables list-columns <table-name> <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		tableName := args[0]
		datasourceId := args[1]

		result, err := client.Datasources.Tables.ListColumns(ctx, tableName, datasourceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tablesCmd.AddCommand(tablesListColumnsCmd)

}
