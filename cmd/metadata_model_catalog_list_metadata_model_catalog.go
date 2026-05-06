package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	metadataModelCatalogListMetadataModelCatalogScope string
)

var metadataModelCatalogListMetadataModelCatalogCmd = &cobra.Command{
	Use:   "list",
	Short: "List Metadata Model Catalog",
	Long:  `List Metadata Model Catalog`,
	Example: "meibel metadata-model-catalog list --scope=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := &sdk.ListMetadataModelCatalogOptions{}
		if metadataModelCatalogListMetadataModelCatalogScope != "" {
			opts.Scope = &metadataModelCatalogListMetadataModelCatalogScope
		}

		result, err := client.MetadataModelCatalog.ListMetadataModelCatalog(ctx, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	metadataModelCatalogCmd.AddCommand(metadataModelCatalogListMetadataModelCatalogCmd)

	metadataModelCatalogListMetadataModelCatalogCmd.Flags().StringVarP(&metadataModelCatalogListMetadataModelCatalogScope, "scope", "", "", "The scope parameter")
}
