package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var (
	dataelementsDeleteDataElementForce bool
)

var dataelementsDeleteDataElementCmd = &cobra.Command{
	Use:   "delete-data-element <datasource-id> <data-element-id>",
	Short: "Delete Data Element",
	Long:  `Delete Data Element

Arguments:
  datasource-id: required
  data-element-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources dataelements delete-data-element <datasource-id> <data-element-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		dataElementId := args[1]

		if !dataelementsDeleteDataElementForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		result, err := client.Datasources.Dataelements.DeleteDataElement(ctx, datasourceId, dataElementId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	dataelementsCmd.AddCommand(dataelementsDeleteDataElementCmd)

	dataelementsDeleteDataElementCmd.Flags().BoolVarP(&dataelementsDeleteDataElementForce, "force", "f", false, "skip confirmation prompt")
}
