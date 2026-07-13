package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	documentsListDeepTransformsOffset int64
	documentsListDeepTransformsLimit int64
)

var documentsListDeepTransformsCmd = &cobra.Command{
	Use:   "list-deep-transforms",
	Short: "List deep-transform jobs",
	Long:  `List the calling customer's deep-transform jobs, newest first. Scoped to the customer (and project, when a project header is set). Paginated via 'offset'/'limit'.`,
	Example: "meibel documents list-deep-transforms --offset=<value> --limit=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := &sdk.DocumentsListDeepTransformsOptions{}
		if documentsListDeepTransformsOffset != 0 {
			opts.Offset = &documentsListDeepTransformsOffset
		}
		if documentsListDeepTransformsLimit != 0 {
			opts.Limit = &documentsListDeepTransformsLimit
		}

		iter := client.Documents.ListDeepTransforms(ctx, opts)

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
	documentsCmd.AddCommand(documentsListDeepTransformsCmd)

	documentsListDeepTransformsCmd.Flags().Int64VarP(&documentsListDeepTransformsOffset, "offset", "", 0, "Number of jobs to skip")
	documentsListDeepTransformsCmd.Flags().Int64VarP(&documentsListDeepTransformsLimit, "limit", "", 50, "Maximum number of jobs to return")
}
