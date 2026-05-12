package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var (
	datasourcesDeleteForce bool
)

var datasourcesDeleteCmd = &cobra.Command{
	Use:   "delete <datasource-id>",
	Short: "Delete Datasource",
	Long:  `Delete Datasource

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources delete <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		if !datasourcesDeleteForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		result, err := client.Datasources.Delete(ctx, datasourceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesCmd.AddCommand(datasourcesDeleteCmd)

	datasourcesDeleteCmd.Flags().BoolVarP(&datasourcesDeleteForce, "force", "f", false, "skip confirmation prompt")
}
