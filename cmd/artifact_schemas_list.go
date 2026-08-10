package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/meibel-ai/meibel-cli/internal/tui"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	artifactSchemasListOffset int64
	artifactSchemasListLimit string
	artifactSchemasListSortBy string
	artifactSchemasListSortOrder string
)

var artifactSchemasListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Artifact Schemas",
	Long:  `List Artifact Schemas`,
	Example: "meibel artifact-schemas list --offset=<value> --limit=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := &sdk.ArtifactSchemasListOptions{}
		if artifactSchemasListOffset != 0 {
			opts.Offset = &artifactSchemasListOffset
		}
		if artifactSchemasListLimit != "" {
			opts.Limit = &artifactSchemasListLimit
		}
		if artifactSchemasListSortBy != "" {
			opts.SortBy = &artifactSchemasListSortBy
		}
		if artifactSchemasListSortOrder != "" {
			opts.SortOrder = &artifactSchemasListSortOrder
		}

		iter := client.ArtifactSchemas.List(ctx, opts)

		sp := tui.StartSpinner("List Artifact Schemas")
		var items []interface{}
		for iter.Next(ctx) {
			items = append(items, iter.Item())
		}
		sp.Stop()
		if err := iter.Err(); err != nil {
			return err
		}

		return output.Print(items)
	},
}

func init() {
	artifactSchemasCmd.AddCommand(artifactSchemasListCmd)

	artifactSchemasListCmd.Flags().Int64VarP(&artifactSchemasListOffset, "offset", "", 0, "Number of items to skip")
	artifactSchemasListCmd.Flags().StringVarP(&artifactSchemasListLimit, "limit", "", "", "Maximum number of items to return")
	artifactSchemasListCmd.Flags().StringVarP(&artifactSchemasListSortBy, "sort-by", "", "", "Field to sort by: created_at, name, display_name")
	artifactSchemasListCmd.Flags().StringVarP(&artifactSchemasListSortOrder, "sort-order", "", "", "Sort order: asc or desc")
}
