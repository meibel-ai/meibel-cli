package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	batchDefinitionsListVersionsOffset int64
	batchDefinitionsListVersionsLimit string
)

var batchDefinitionsListVersionsCmd = &cobra.Command{
	Use:   "list-versions <definition-id>",
	Short: "List Batch Definition Versions",
	Long:  `List Batch Definition Versions

Arguments:
  definition-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel batch-definitions list-versions <definition-id> --offset=<value> --limit=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		definitionId := args[0]

		opts := &sdk.BatchDefinitionsListVersionsOptions{}
		if batchDefinitionsListVersionsOffset != 0 {
			opts.Offset = &batchDefinitionsListVersionsOffset
		}
		if batchDefinitionsListVersionsLimit != "" {
			opts.Limit = &batchDefinitionsListVersionsLimit
		}

		iter := client.BatchDefinitions.ListVersions(ctx, definitionId, opts)

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
	batchDefinitionsCmd.AddCommand(batchDefinitionsListVersionsCmd)

	batchDefinitionsListVersionsCmd.Flags().Int64VarP(&batchDefinitionsListVersionsOffset, "offset", "", 0, "The offset parameter")
	batchDefinitionsListVersionsCmd.Flags().StringVarP(&batchDefinitionsListVersionsLimit, "limit", "", "", "The limit parameter")
}
