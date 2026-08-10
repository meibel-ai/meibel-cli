package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/meibel-ai/meibel-cli/internal/tui"
)

var executionPoliciesGetCmd = &cobra.Command{
	Use:   "get <policy-id>",
	Short: "Get Execution Policy",
	Long:  `Get Execution Policy

Arguments:
  policy-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel execution-policies get <policy-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		policyId := args[0]

		sp := tui.StartSpinner("Get Execution Policy")
		result, err := client.ExecutionPolicies.Get(ctx, policyId)
		sp.Stop()
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	executionPoliciesCmd.AddCommand(executionPoliciesGetCmd)

}
