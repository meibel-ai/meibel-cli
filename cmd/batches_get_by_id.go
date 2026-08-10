package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/meibel-ai/meibel-cli/internal/tui"
)

var batchesGetByIdCmd = &cobra.Command{
	Use:   "get-by-id <definition-id>",
	Short: "Get Batch Definition By Id",
	Long:  `Get Batch Definition By Id

Arguments:
  definition-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel batches get-by-id <definition-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		definitionId := args[0]

		sp := tui.StartSpinner("Get Batch Definition By Id")
		result, err := client.Batches.GetById(ctx, definitionId)
		sp.Stop()
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	batchesCmd.AddCommand(batchesGetByIdCmd)

}
