package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	tableDescriptionsListTablesIncludeColumns bool
)

var tableDescriptionsListTablesCmd = &cobra.Command{
	Use:   "list-tables <datasource-id>",
	Short: "List Tables",
	Long:  `List Tables

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources table-descriptions list-tables <datasource-id> --include-columns=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		opts := &sdk.ListTablesOptions{}
		if tableDescriptionsListTablesIncludeColumns {
			opts.IncludeColumns = &tableDescriptionsListTablesIncludeColumns
		}

		result, err := client.Datasources.TableDescriptions.ListTables(ctx, datasourceId, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tableDescriptionsCmd.AddCommand(tableDescriptionsListTablesCmd)

	tableDescriptionsListTablesCmd.Flags().BoolVarP(&tableDescriptionsListTablesIncludeColumns, "include-columns", "", false, "Include columns for each table")
}
