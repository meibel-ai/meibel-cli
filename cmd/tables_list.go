package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	tablesListIncludeColumns bool
)

var tablesListCmd = &cobra.Command{
	Use:   "list <datasource-id>",
	Short: "List Tables",
	Long:  `List Tables

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources tables list <datasource-id> --include-columns=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		opts := &sdk.TablesListOptions{}
		if tablesListIncludeColumns {
			opts.IncludeColumns = &tablesListIncludeColumns
		}

		result, err := client.Datasources.Tables.List(ctx, datasourceId, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tablesCmd.AddCommand(tablesListCmd)

	tablesListCmd.Flags().BoolVarP(&tablesListIncludeColumns, "include-columns", "", false, "Include columns for each table")
}
