package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	datasourcesGetDatasourceIncludeTables bool
)

var datasourcesGetDatasourceCmd = &cobra.Command{
	Use:   "get <datasource-id>",
	Short: "Get Datasource",
	Long:  `Get Datasource

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources get <datasource-id> --include-tables=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		opts := &sdk.GetDatasourceOptions{}
		if datasourcesGetDatasourceIncludeTables {
			opts.IncludeTables = &datasourcesGetDatasourceIncludeTables
		}

		result, err := client.Datasources.GetDatasource(ctx, datasourceId, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesCmd.AddCommand(datasourcesGetDatasourceCmd)

	datasourcesGetDatasourceCmd.Flags().BoolVarP(&datasourcesGetDatasourceIncludeTables, "include-tables", "", false, "Include table and column details (structured datasources only)")
}
