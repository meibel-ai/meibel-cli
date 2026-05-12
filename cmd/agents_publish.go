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
	agentsPublishOverrideDraft bool
	agentsPublishData string
	agentsPublishInteractive bool
)

var agentsPublishCmd = &cobra.Command{
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

		if agentsPublishData != "" {
			if err := json.Unmarshal([]byte(agentsPublishData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if agentsPublishInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		opts := &sdk.AgentsPublishOptions{}
		if agentsPublishOverrideDraft {
			opts.OverrideDraft = &agentsPublishOverrideDraft
		}

		result, err := client.Agents.Publish(ctx, agentId, body, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	agentsCmd.AddCommand(agentsPublishCmd)

	agentsPublishCmd.Flags().BoolVarP(&agentsPublishOverrideDraft, "override-draft", "", false, "Bypass draft head validation and publish any version directly")
	agentsPublishCmd.Flags().StringVarP(&agentsPublishData, "data", "d", "", "JSON data for the request body")
	agentsPublishCmd.Flags().BoolVarP(&agentsPublishInteractive, "interactive", "i", false, "use interactive form input")
}
