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
	agentsCreateAgentData string
	agentsCreateAgentInteractive bool
)

var agentsCreateAgentCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Agent",
	Long:  `Create Agent`,
	Example: "meibel agents create",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		var body sdk.CreateAgentDefinitionRequest

		if agentsCreateAgentData != "" {
			if err := json.Unmarshal([]byte(agentsCreateAgentData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if agentsCreateAgentInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("DisplayName").Description("Human-readable name of the agent (letters, numbers, and spaces only). Converted to kebab-case internally.").Value(&body.DisplayName),
					huh.NewInput().Title("Instructions").Description("System prompt/instructions for the agent").Value(&body.Instructions),
					huh.NewInput().Title("AdditionalProperties").Description(""),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.Agents.CreateAgent(ctx, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	agentsCmd.AddCommand(agentsCreateAgentCmd)

	agentsCreateAgentCmd.Flags().StringVarP(&agentsCreateAgentData, "data", "d", "", "JSON data for the request body")
	agentsCreateAgentCmd.Flags().BoolVarP(&agentsCreateAgentInteractive, "interactive", "i", false, "use interactive form input")
}
