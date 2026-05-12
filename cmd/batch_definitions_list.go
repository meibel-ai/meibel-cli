package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	batchDefinitionsListOffset int64
	batchDefinitionsListLimit int64
)

var batchDefinitionsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Batch Definitions",
	Long:  `List Batch Definitions`,
	Example: "meibel batch-definitions list --offset=<value> --limit=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := &sdk.BatchDefinitionsListOptions{}
		if batchDefinitionsListOffset != 0 {
			opts.Offset = &batchDefinitionsListOffset
		}
		if batchDefinitionsListLimit != 0 {
			opts.Limit = &batchDefinitionsListLimit
		}

		iter := client.BatchDefinitions.List(ctx, opts)

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
	batchDefinitionsCmd.AddCommand(batchDefinitionsListCmd)

	batchDefinitionsListCmd.Flags().Int64VarP(&batchDefinitionsListOffset, "offset", "", 0, "The offset parameter")
	batchDefinitionsListCmd.Flags().Int64VarP(&batchDefinitionsListLimit, "limit", "", 100, "The limit parameter")
}
