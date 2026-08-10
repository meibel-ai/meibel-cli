package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/tui"
)

var (
	executionPoliciesDeleteForce bool
)

var executionPoliciesDeleteCmd = &cobra.Command{
	Use:   "delete <policy-id>",
	Short: "Delete Execution Policy",
	Long:  `Delete Execution Policy

Arguments:
  policy-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel execution-policies delete <policy-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		policyId := args[0]

		if !executionPoliciesDeleteForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		sp := tui.StartSpinner("Delete Execution Policy")
		err := client.ExecutionPolicies.Delete(ctx, policyId)
		sp.Stop()
		if err != nil {
			return err
		}

		fmt.Println("Success")
		return nil
	},
}

func init() {
	executionPoliciesCmd.AddCommand(executionPoliciesDeleteCmd)

	executionPoliciesDeleteCmd.Flags().BoolVarP(&executionPoliciesDeleteForce, "force", "f", false, "skip confirmation prompt")
}
