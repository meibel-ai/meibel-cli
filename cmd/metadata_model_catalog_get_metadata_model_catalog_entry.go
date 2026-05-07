package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var metadataModelCatalogGetMetadataModelCatalogEntryCmd = &cobra.Command{
	Use:   "get-entry <model-id>",
	Short: "Get Metadata Model Catalog Entry",
	Long:  `Get Metadata Model Catalog Entry

Arguments:
  model-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel metadata-model-catalog get-entry <model-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		modelId := args[0]

		result, err := client.MetadataModelCatalog.GetMetadataModelCatalogEntry(ctx, modelId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	metadataModelCatalogCmd.AddCommand(metadataModelCatalogGetMetadataModelCatalogEntryCmd)

}
