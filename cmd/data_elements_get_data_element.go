package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel/internal/output"
)

var dataElementsGetDataElementCmd = &cobra.Command{
	Use:   "get <datasource-id> <data-element-id>",
	Short: "Get Data Element",
	Long:  `Get Data Element

Arguments:
  datasource-id: required
  data-element-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel data-elements get <datasource-id> <data-element-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		dataElementId := args[1]

		result, err := client.DataElements.GetDataElement(ctx, datasourceId, dataElementId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	dataElementsCmd.AddCommand(dataElementsGetDataElementCmd)

}
