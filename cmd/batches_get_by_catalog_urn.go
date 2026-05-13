package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var (
	batchesGetByCatalogUrnCatalogUrn string
)

var batchesGetByCatalogUrnCmd = &cobra.Command{
	Use:   "get-by-catalog-urn",
	Short: "Get Batch Definition By Catalog Urn",
	Long:  `Get Batch Definition By Catalog Urn`,
	Example: "meibel batches get-by-catalog-urn",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		result, err := client.Batches.GetByCatalogUrn(ctx, batchesGetByCatalogUrnCatalogUrn)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	batchesCmd.AddCommand(batchesGetByCatalogUrnCmd)

	batchesGetByCatalogUrnCmd.Flags().StringVarP(&batchesGetByCatalogUrnCatalogUrn, "catalog-urn", "", "", "urn:meibel:batch-definition:{customer}:{project}:{id}")
	batchesGetByCatalogUrnCmd.MarkFlagRequired("catalog-urn")
}
