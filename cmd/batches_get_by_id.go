package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
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

		result, err := client.Batches.GetById(ctx, definitionId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	batchesCmd.AddCommand(batchesGetByIdCmd)

}
