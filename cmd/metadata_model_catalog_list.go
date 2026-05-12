package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	metadataModelCatalogListScope string
)

var metadataModelCatalogListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Metadata Model Catalog",
	Long:  `List Metadata Model Catalog`,
	Example: "meibel metadata-model-catalog list --scope=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := &sdk.MetadataModelCatalogListOptions{}
		if metadataModelCatalogListScope != "" {
			opts.Scope = &metadataModelCatalogListScope
		}

		result, err := client.MetadataModelCatalog.List(ctx, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	metadataModelCatalogCmd.AddCommand(metadataModelCatalogListCmd)

	metadataModelCatalogListCmd.Flags().StringVarP(&metadataModelCatalogListScope, "scope", "", "", "The scope parameter")
}
