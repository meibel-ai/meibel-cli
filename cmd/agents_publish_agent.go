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
	agentsPublishAgentOverrideDraft bool
	agentsPublishAgentData string
	agentsPublishAgentInteractive bool
)

var agentsPublishAgentCmd = &cobra.Command{
	Use:   "publish <agent-id>",
	Short: "Publish Agent",
	Long:  `Publish Agent

Arguments:
  agent-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel agents publish <agent-id> --override-draft=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		agentId := args[0]

		var body sdk.PublishAgentDefinitionRequest

		if agentsPublishAgentData != "" {
			if err := json.Unmarshal([]byte(agentsPublishAgentData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if agentsPublishAgentInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("CommitMessage").Description("User-provided description of what changed in this version").Value(&body.CommitMessage),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		opts := &sdk.PublishAgentOptions{}
		if agentsPublishAgentOverrideDraft {
			opts.OverrideDraft = &agentsPublishAgentOverrideDraft
		}

		result, err := client.Agents.PublishAgent(ctx, agentId, body, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	agentsCmd.AddCommand(agentsPublishAgentCmd)

	agentsPublishAgentCmd.Flags().BoolVarP(&agentsPublishAgentOverrideDraft, "override-draft", "", false, "Bypass draft head validation and publish any version directly")
	agentsPublishAgentCmd.Flags().StringVarP(&agentsPublishAgentData, "data", "d", "", "JSON data for the request body")
	agentsPublishAgentCmd.Flags().BoolVarP(&agentsPublishAgentInteractive, "interactive", "i", false, "use interactive form input")
}
