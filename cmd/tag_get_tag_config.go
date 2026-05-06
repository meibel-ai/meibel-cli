package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var tagGetTagConfigCmd = &cobra.Command{
	Use:   "get-config <datasource-id>",
	Short: "Get Tag Config",
	Long:  `Get Tag Config

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources tag get-config <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		result, err := client.Datasources.Tag.GetTagConfig(ctx, datasourceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tagCmd.AddCommand(tagGetTagConfigCmd)

}
