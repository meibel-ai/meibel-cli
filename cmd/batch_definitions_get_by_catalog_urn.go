package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var (
	batchDefinitionsGetByCatalogUrnCatalogUrn string
)

var batchDefinitionsGetByCatalogUrnCmd = &cobra.Command{
	Use:   "get-by-catalog-urn",
	Short: "Get Batch Definition By Catalog Urn",
	Long:  `Get Batch Definition By Catalog Urn`,
	Example: "meibel batch-definitions get-by-catalog-urn",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		result, err := client.BatchDefinitions.GetByCatalogUrn(ctx, batchDefinitionsGetByCatalogUrnCatalogUrn)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	batchDefinitionsCmd.AddCommand(batchDefinitionsGetByCatalogUrnCmd)

	batchDefinitionsGetByCatalogUrnCmd.Flags().StringVarP(&batchDefinitionsGetByCatalogUrnCatalogUrn, "catalog-urn", "", "", "urn:meibel:batch-definition:{customer}:{project}:{id}")
	batchDefinitionsGetByCatalogUrnCmd.MarkFlagRequired("catalog-urn")
}
