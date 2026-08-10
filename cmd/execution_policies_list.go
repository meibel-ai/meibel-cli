package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/meibel-ai/meibel-cli/internal/tui"
)

var executionPoliciesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Execution Policies",
	Long:  `List Execution Policies`,
	Example: "meibel execution-policies list",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		sp := tui.StartSpinner("List Execution Policies")
		result, err := client.ExecutionPolicies.List(ctx)
		sp.Stop()
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	executionPoliciesCmd.AddCommand(executionPoliciesListCmd)

}
