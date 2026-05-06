package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
	"golang.org/x/term"
	"github.com/meibel-ai/meibel/internal/tui"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	agentsSessionsSendChatMessageStreamData string
	agentsSessionsSendChatMessageStreamInteractive bool
)

var agentsSessionsSendChatMessageStreamCmd = &cobra.Command{
	Use:   "send-chat-message-stream <session-id>",
	Short: "Send a chat message and stream the response via SSE",
	Long:  `Send a chat message and stream the response via SSE

Arguments:
  session-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel agents agents-sessions send-chat-message-stream <session-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		sessionId := args[0]

		var body sdk.ChatMessageRequest

		if agentsSessionsSendChatMessageStreamData != "" {
			if err := json.Unmarshal([]byte(agentsSessionsSendChatMessageStreamData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if agentsSessionsSendChatMessageStreamInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		// Set up signal handling for graceful shutdown
		ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
		defer cancel()

		stream, err := client.Agents.Sessions.SendChatMessageStream(ctx, sessionId, body)
		if err != nil {
			return err
		}
		defer stream.Close()

		return tui.StreamEvents(ctx, stream)
	},
}

func init() {
	agentsSessionsCmd.AddCommand(agentsSessionsSendChatMessageStreamCmd)

	agentsSessionsSendChatMessageStreamCmd.Flags().StringVarP(&agentsSessionsSendChatMessageStreamData, "data", "d", "", "JSON data for the request body")
	agentsSessionsSendChatMessageStreamCmd.Flags().BoolVarP(&agentsSessionsSendChatMessageStreamInteractive, "interactive", "i", false, "use interactive form input")
}
