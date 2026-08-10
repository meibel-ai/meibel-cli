package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/meibel-ai/meibel-cli/internal/tui"
)

var metadataModelCatalogGetEntryCmd = &cobra.Command{
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

		sp := tui.StartSpinner("Get Metadata Model Catalog Entry")
		result, err := client.MetadataModelCatalog.GetEntry(ctx, modelId)
		sp.Stop()
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	metadataModelCatalogCmd.AddCommand(metadataModelCatalogGetEntryCmd)

}
