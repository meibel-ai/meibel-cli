package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var (
	contentDeleteDatasourceContentForce bool
)

var contentDeleteDatasourceContentCmd = &cobra.Command{
	Use:   "delete-datasource <datasource-id> <path>",
	Short: "Delete content",
	Long:  `Delete a file or directory from the datasource

Arguments:
  datasource-id: required
  path: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources content delete-datasource <datasource-id> <path>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		path := args[1]

		if !contentDeleteDatasourceContentForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		result, err := client.Datasources.Content.DeleteDatasourceContent(ctx, datasourceId, path)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	contentCmd.AddCommand(contentDeleteDatasourceContentCmd)

	contentDeleteDatasourceContentCmd.Flags().BoolVarP(&contentDeleteDatasourceContentForce, "force", "f", false, "skip confirmation prompt")
}
