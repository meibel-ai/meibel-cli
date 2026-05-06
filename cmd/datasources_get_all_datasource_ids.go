package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var datasourcesGetAllDatasourceIdsCmd = &cobra.Command{
	Use:   "get-all-ids",
	Short: "Get All Datasource Ids",
	Long:  `Get All Datasource Ids`,
	Example: "meibel datasources get-all-ids",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		result, err := client.Datasources.GetAllDatasourceIds(ctx)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesCmd.AddCommand(datasourcesGetAllDatasourceIdsCmd)

}
