package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	agentsListOffset int64
	agentsListLimit int64
	agentsListSortBy string
	agentsListSortOrder string
	agentsListPublishedOnly bool
	agentsListDatasourceId string
	agentsListArtifactSchemaId string
)

var agentsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Agents",
	Long:  `List Agents`,
	Example: "meibel agents list --offset=<value> --limit=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := &sdk.AgentsListOptions{}
		if agentsListOffset != 0 {
			opts.Offset = &agentsListOffset
		}
		if agentsListLimit != 0 {
			opts.Limit = &agentsListLimit
		}
		if agentsListSortBy != "" {
			opts.SortBy = &agentsListSortBy
		}
		if agentsListSortOrder != "" {
			opts.SortOrder = &agentsListSortOrder
		}
		if agentsListPublishedOnly {
			opts.PublishedOnly = &agentsListPublishedOnly
		}
		if agentsListDatasourceId != "" {
			opts.DatasourceId = &agentsListDatasourceId
		}
		if agentsListArtifactSchemaId != "" {
			opts.ArtifactSchemaId = &agentsListArtifactSchemaId
		}

		iter := client.Agents.List(ctx, opts)

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
	agentsCmd.AddCommand(agentsListCmd)

	agentsListCmd.Flags().Int64VarP(&agentsListOffset, "offset", "", 0, "Number of items to skip")
	agentsListCmd.Flags().Int64VarP(&agentsListLimit, "limit", "", 20, "Maximum number of items to return")
	agentsListCmd.Flags().StringVarP(&agentsListSortBy, "sort-by", "", "created_at", "Field to sort by: created_at, name, display_name")
	agentsListCmd.Flags().StringVarP(&agentsListSortOrder, "sort-order", "", "desc", "Sort order: asc or desc")
	agentsListCmd.Flags().BoolVarP(&agentsListPublishedOnly, "published-only", "", false, "If true, return only published agents (latest published version per name)")
	agentsListCmd.Flags().StringVarP(&agentsListDatasourceId, "datasource-id", "", "", "Return only agents whose latest version uses this datasource ID")
	agentsListCmd.Flags().StringVarP(&agentsListArtifactSchemaId, "artifact-schema-id", "", "", "Return only agents whose latest version produces this artifact (catalog URN)")
}
