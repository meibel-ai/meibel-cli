package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
	"golang.org/x/term"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	executionPoliciesUpdateData string
	executionPoliciesUpdateInteractive bool
)

var executionPoliciesUpdateCmd = &cobra.Command{
	Use:   "update <policy-id>",
	Short: "Update Execution Policy",
	Long:  `Update Execution Policy

Arguments:
  policy-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel execution-policies update <policy-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		policyId := args[0]

		var body sdk.UpdateExecutionPolicyRequest

		if executionPoliciesUpdateData != "" {
			if err := json.Unmarshal([]byte(executionPoliciesUpdateData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if executionPoliciesUpdateInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.ExecutionPolicies.Update(ctx, policyId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	executionPoliciesCmd.AddCommand(executionPoliciesUpdateCmd)

	executionPoliciesUpdateCmd.Flags().StringVarP(&executionPoliciesUpdateData, "data", "d", "", "JSON data for the request body")
	executionPoliciesUpdateCmd.Flags().BoolVarP(&executionPoliciesUpdateInteractive, "interactive", "i", false, "use interactive form input")
}
