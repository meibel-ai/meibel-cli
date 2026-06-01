package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var dataElementsGetCmd = &cobra.Command{
	Use:   "get <data-element-id> <datasource-id>",
	Short: "Get Data Element",
	Long:  `Get Data Element

Arguments:
  data-element-id: required
  datasource-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources data-elements get <data-element-id> <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		dataElementId := args[0]
		datasourceId := args[1]

		result, err := client.Datasources.DataElements.Get(ctx, dataElementId, datasourceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	dataElementsCmd.AddCommand(dataElementsGetCmd)

}
