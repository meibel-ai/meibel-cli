package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var executionPoliciesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Execution Policies",
	Long:  `List Execution Policies`,
	Example: "meibel execution-policies list",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		result, err := client.ExecutionPolicies.List(ctx)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	executionPoliciesCmd.AddCommand(executionPoliciesListCmd)

}
