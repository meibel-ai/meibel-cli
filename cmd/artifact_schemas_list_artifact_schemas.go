package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	artifactSchemasListArtifactSchemasOffset int64
	artifactSchemasListArtifactSchemasLimit string
	artifactSchemasListArtifactSchemasSortBy string
	artifactSchemasListArtifactSchemasSortOrder string
)

var artifactSchemasListArtifactSchemasCmd = &cobra.Command{
	Use:   "list",
	Short: "List Artifact Schemas",
	Long:  `List Artifact Schemas`,
	Example: "meibel artifact-schemas list --offset=<value> --limit=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := &sdk.ListArtifactSchemasOptions{}
		if artifactSchemasListArtifactSchemasOffset != 0 {
			opts.Offset = &artifactSchemasListArtifactSchemasOffset
		}
		if artifactSchemasListArtifactSchemasLimit != "" {
			opts.Limit = &artifactSchemasListArtifactSchemasLimit
		}
		if artifactSchemasListArtifactSchemasSortBy != "" {
			opts.SortBy = &artifactSchemasListArtifactSchemasSortBy
		}
		if artifactSchemasListArtifactSchemasSortOrder != "" {
			opts.SortOrder = &artifactSchemasListArtifactSchemasSortOrder
		}

		iter := client.ArtifactSchemas.ListArtifactSchemas(ctx, opts)

		var items []interface{}
		for iter.Next(ctx) {
			items = append(items, iter.Item())
		}
		if err := iter.Err(); err != nil {
			return err
		}

		return output.Print(items)
	},
}

func init() {
	artifactSchemasCmd.AddCommand(artifactSchemasListArtifactSchemasCmd)

	artifactSchemasListArtifactSchemasCmd.Flags().Int64VarP(&artifactSchemasListArtifactSchemasOffset, "offset", "", 0, "Number of items to skip")
	artifactSchemasListArtifactSchemasCmd.Flags().StringVarP(&artifactSchemasListArtifactSchemasLimit, "limit", "", "", "Maximum number of items to return")
	artifactSchemasListArtifactSchemasCmd.Flags().StringVarP(&artifactSchemasListArtifactSchemasSortBy, "sort-by", "", "", "Field to sort by: created_at, name, display_name")
	artifactSchemasListArtifactSchemasCmd.Flags().StringVarP(&artifactSchemasListArtifactSchemasSortOrder, "sort-order", "", "", "Sort order: asc or desc")
}
