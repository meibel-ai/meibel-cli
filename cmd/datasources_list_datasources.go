package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel/internal/output"
)

var datasourcesListDatasourcesCmd = &cobra.Command{
	Use:   "list",
	Short: "List Datasources",
	Long:  `List Datasources`,
	Example: "meibel datasources list",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		result, err := client.Datasources.ListDatasources(ctx)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesCmd.AddCommand(datasourcesListDatasourcesCmd)

}
