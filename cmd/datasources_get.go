package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	datasourcesGetIncludeTables bool
)

var datasourcesGetCmd = &cobra.Command{
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

		opts := &sdk.DatasourcesGetOptions{}
		if datasourcesGetIncludeTables {
			opts.IncludeTables = &datasourcesGetIncludeTables
		}

		result, err := client.Datasources.Get(ctx, datasourceId, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesCmd.AddCommand(datasourcesGetCmd)

	datasourcesGetCmd.Flags().BoolVarP(&datasourcesGetIncludeTables, "include-tables", "", false, "Include table and column details (structured datasources only)")
}
