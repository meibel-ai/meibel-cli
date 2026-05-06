package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var (
	tagDeleteTagColumnInfoForce bool
)

var tagDeleteTagColumnInfoCmd = &cobra.Command{
	Use:   "delete-column-info <datasource-id> <table-name> <column-name>",
	Short: "Delete Tag Column Info",
	Long:  `Delete Tag Column Info

Arguments:
  datasource-id: required
  table-name: required
  column-name: required`,
	Args:  cobra.ExactArgs(3),
	Example: "meibel datasources tag delete-column-info <datasource-id> <table-name> <column-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		tableName := args[1]
		columnName := args[2]

		if !tagDeleteTagColumnInfoForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		result, err := client.Datasources.Tag.DeleteTagColumnInfo(ctx, datasourceId, tableName, columnName)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tagCmd.AddCommand(tagDeleteTagColumnInfoCmd)

	tagDeleteTagColumnInfoCmd.Flags().BoolVarP(&tagDeleteTagColumnInfoForce, "force", "f", false, "skip confirmation prompt")
}
