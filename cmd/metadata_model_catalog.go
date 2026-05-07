package cmd

import (
	"github.com/spf13/cobra"
)

var metadataModelCatalogCmd = &cobra.Command{
	Use:   "metadata-model-catalog",
	Short: "Manage Metadata Model Catalog",
	Long:  `Commands for managing Metadata Model Catalog resources.`,
}

func init() {
	rootCmd.AddCommand(metadataModelCatalogCmd)
}
