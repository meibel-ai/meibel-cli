package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/meibel-ai/meibel-cli/internal/tui"
)

var batchesExecuteCmd = &cobra.Command{
	Use:   "execute <definition-id>",
	Short: "Execute Batch Definition",
	Long:  `Execute Batch Definition

Arguments:
  definition-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel batches execute <definition-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		definitionId := args[0]

		sp := tui.StartSpinner("Execute Batch Definition")
		result, err := client.Batches.Execute(ctx, definitionId)
		sp.Stop()
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	batchesCmd.AddCommand(batchesExecuteCmd)

}
