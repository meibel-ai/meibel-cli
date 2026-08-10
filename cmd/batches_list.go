package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/meibel-ai/meibel-cli/internal/tui"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	batchesListOffset int64
	batchesListLimit int64
)

var batchesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Batch Definitions",
	Long:  `List Batch Definitions`,
	Example: "meibel batches list --offset=<value> --limit=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := &sdk.BatchesListOptions{}
		if batchesListOffset != 0 {
			opts.Offset = &batchesListOffset
		}
		if batchesListLimit != 0 {
			opts.Limit = &batchesListLimit
		}

		sp := tui.StartSpinner("List Batch Definitions")
		result, err := client.Batches.List(ctx, opts)
		sp.Stop()
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	batchesCmd.AddCommand(batchesListCmd)

	batchesListCmd.Flags().Int64VarP(&batchesListOffset, "offset", "", 0, "The offset parameter")
	batchesListCmd.Flags().Int64VarP(&batchesListLimit, "limit", "", 100, "The limit parameter")
}
