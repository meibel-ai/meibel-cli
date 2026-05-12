package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var batchDefinitionsExecuteCmd = &cobra.Command{
	Use:   "execute <definition-id>",
	Short: "Execute Batch Definition",
	Long:  `Execute Batch Definition

Arguments:
  definition-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel batch-definitions execute <definition-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		definitionId := args[0]

		result, err := client.BatchDefinitions.Execute(ctx, definitionId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	batchDefinitionsCmd.AddCommand(batchDefinitionsExecuteCmd)

}
