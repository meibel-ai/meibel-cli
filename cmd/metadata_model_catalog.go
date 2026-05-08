package cmd

import (
	"github.com/spf13/cobra"
)

var metadataModelCatalogCmd = &cobra.Command{
	Use:   "metadata-model-catalog",
	Short: "Manage MetadataModelCatalog",
	Long:  `Commands for managing MetadataModelCatalog resources.`,
}

func init() {
	rootCmd.AddCommand(metadataModelCatalogCmd)
}
