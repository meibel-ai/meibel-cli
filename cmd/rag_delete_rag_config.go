package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var (
	ragDeleteRagConfigForce bool
)

var ragDeleteRagConfigCmd = &cobra.Command{
	Use:   "delete-config <datasource-id>",
	Short: "Delete Rag Config",
	Long:  `Delete Rag Config

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources rag delete-config <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		if !ragDeleteRagConfigForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		result, err := client.Datasources.Rag.DeleteRagConfig(ctx, datasourceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	ragCmd.AddCommand(ragDeleteRagConfigCmd)

	ragDeleteRagConfigCmd.Flags().BoolVarP(&ragDeleteRagConfigForce, "force", "f", false, "skip confirmation prompt")
}
