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
	sessionsSendChatMessageStreamData string
	sessionsSendChatMessageStreamInteractive bool
)

var sessionsSendChatMessageStreamCmd = &cobra.Command{
	Use:   "send-chat-message-stream <session-id>",
	Short: "Send a chat message and stream the response via SSE",
	Long:  `Send a chat message and stream the response via SSE

Arguments:
  session-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel sessions send-chat-message-stream <session-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		sessionId := args[0]

		var body sdk.ChatMessageRequest

		if sessionsSendChatMessageStreamData != "" {
			if err := json.Unmarshal([]byte(sessionsSendChatMessageStreamData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if sessionsSendChatMessageStreamInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		stream, err := client.Sessions.SendChatMessageStream(ctx, sessionId, body)
		if err != nil {
			return err
		}
		defer stream.Close()

		return tui.StreamEvents(ctx, stream)
	},
}

func init() {
	sessionsCmd.AddCommand(sessionsSendChatMessageStreamCmd)

	sessionsSendChatMessageStreamCmd.Flags().StringVarP(&sessionsSendChatMessageStreamData, "data", "d", "", "JSON data for the request body")
	sessionsSendChatMessageStreamCmd.Flags().BoolVarP(&sessionsSendChatMessageStreamInteractive, "interactive", "i", false, "use interactive form input")
}
