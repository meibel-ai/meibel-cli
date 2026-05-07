package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
	"golang.org/x/term"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	agentsUpdateAgentData string
	agentsUpdateAgentInteractive bool
)

var agentsUpdateAgentCmd = &cobra.Command{
	Use:   "update <agent-id>",
	Short: "Update Agent",
	Long:  `Update Agent

Arguments:
  agent-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel agents update <agent-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		agentId := args[0]

		var body sdk.UpdateAgentDefinitionRequest

		if agentsUpdateAgentData != "" {
			if err := json.Unmarshal([]byte(agentsUpdateAgentData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if agentsUpdateAgentInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Agents.UpdateAgent(ctx, agentId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	agentsCmd.AddCommand(agentsUpdateAgentCmd)

	agentsUpdateAgentCmd.Flags().StringVarP(&agentsUpdateAgentData, "data", "d", "", "JSON data for the request body")
	agentsUpdateAgentCmd.Flags().BoolVarP(&agentsUpdateAgentInteractive, "interactive", "i", false, "use interactive form input")
}
