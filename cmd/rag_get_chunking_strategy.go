package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var ragGetChunkingStrategyCmd = &cobra.Command{
	Use:   "get-chunking-strategy <datasource-id>",
	Short: "Get Chunking Strategy",
	Long:  `Get Chunking Strategy

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources rag get-chunking-strategy <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		result, err := client.Datasources.Rag.GetChunkingStrategy(ctx, datasourceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	ragCmd.AddCommand(ragGetChunkingStrategyCmd)

}
