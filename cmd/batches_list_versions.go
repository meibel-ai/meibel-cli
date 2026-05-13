package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	batchesListVersionsOffset int64
	batchesListVersionsLimit string
)

var batchesListVersionsCmd = &cobra.Command{
	Use:   "list-versions <definition-id>",
	Short: "List Batch Definition Versions",
	Long:  `List Batch Definition Versions

Arguments:
  definition-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel batches list-versions <definition-id> --offset=<value> --limit=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		definitionId := args[0]

		opts := &sdk.BatchesListVersionsOptions{}
		if batchesListVersionsOffset != 0 {
			opts.Offset = &batchesListVersionsOffset
		}
		if batchesListVersionsLimit != "" {
			opts.Limit = &batchesListVersionsLimit
		}

		result, err := client.Batches.ListVersions(ctx, definitionId, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	batchesCmd.AddCommand(batchesListVersionsCmd)

	batchesListVersionsCmd.Flags().Int64VarP(&batchesListVersionsOffset, "offset", "", 0, "The offset parameter")
	batchesListVersionsCmd.Flags().StringVarP(&batchesListVersionsLimit, "limit", "", "", "The limit parameter")
}
