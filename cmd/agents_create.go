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
	agentsCreateData string
	agentsCreateInteractive bool
)

var agentsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Agent",
	Long:  `Create Agent`,
	Example: "meibel agents create",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		var body sdk.CreateAgentDefinitionRequest

		if agentsCreateData != "" {
			if err := json.Unmarshal([]byte(agentsCreateData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if agentsCreateInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("DisplayName").Description("Human-readable name of the agent (letters, numbers, and spaces only). Converted to kebab-case internally.").Value(&body.DisplayName),
					huh.NewInput().Title("Instructions").Description("System prompt/instructions for the agent").Value(&body.Instructions),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.Agents.Create(ctx, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	agentsCmd.AddCommand(agentsCreateCmd)

	agentsCreateCmd.Flags().StringVarP(&agentsCreateData, "data", "d", "", "JSON data for the request body")
	agentsCreateCmd.Flags().BoolVarP(&agentsCreateInteractive, "interactive", "i", false, "use interactive form input")
}
