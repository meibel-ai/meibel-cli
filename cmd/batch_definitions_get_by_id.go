package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var batchDefinitionsGetByIdCmd = &cobra.Command{
	Use:   "get-by-id <definition-id>",
	Short: "Get Batch Definition By Id",
	Long:  `Get Batch Definition By Id

Arguments:
  definition-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel batch-definitions get-by-id <definition-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		definitionId := args[0]

		result, err := client.BatchDefinitions.GetById(ctx, definitionId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	batchDefinitionsCmd.AddCommand(batchDefinitionsGetByIdCmd)

}
