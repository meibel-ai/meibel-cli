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
	executionPoliciesCreateData string
	executionPoliciesCreateInteractive bool
)

var executionPoliciesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Execution Policy",
	Long:  `Create Execution Policy`,
	Example: "meibel execution-policies create",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		var body sdk.CreateExecutionPolicyRequest

		if executionPoliciesCreateData != "" {
			if err := json.Unmarshal([]byte(executionPoliciesCreateData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if executionPoliciesCreateInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("Name").Description("Policy name (unique within tenant)").Value(&body.Name),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.ExecutionPolicies.Create(ctx, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	executionPoliciesCmd.AddCommand(executionPoliciesCreateCmd)

	executionPoliciesCreateCmd.Flags().StringVarP(&executionPoliciesCreateData, "data", "d", "", "JSON data for the request body")
	executionPoliciesCreateCmd.Flags().BoolVarP(&executionPoliciesCreateInteractive, "interactive", "i", false, "use interactive form input")
}
