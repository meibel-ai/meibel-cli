package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
	"golang.org/x/term"
	"github.com/meibel-ai/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	agentsSessionsSendChatMessageData string
	agentsSessionsSendChatMessageInteractive bool
)

var agentsSessionsSendChatMessageCmd = &cobra.Command{
	Use:   "send-chat-message <session-id>",
	Short: "Send Chat Message",
	Long:  `Send Chat Message

Arguments:
  session-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel agents agents-sessions send-chat-message <session-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		sessionId := args[0]

		var body sdk.ChatMessageRequest

		if agentsSessionsSendChatMessageData != "" {
			if err := json.Unmarshal([]byte(agentsSessionsSendChatMessageData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if agentsSessionsSendChatMessageInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("UserMessage").Description("The user's chat message").Value(&body.UserMessage),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.Agents.Sessions.SendChatMessage(ctx, sessionId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	agentsSessionsCmd.AddCommand(agentsSessionsSendChatMessageCmd)

	agentsSessionsSendChatMessageCmd.Flags().StringVarP(&agentsSessionsSendChatMessageData, "data", "d", "", "JSON data for the request body")
	agentsSessionsSendChatMessageCmd.Flags().BoolVarP(&agentsSessionsSendChatMessageInteractive, "interactive", "i", false, "use interactive form input")
}
