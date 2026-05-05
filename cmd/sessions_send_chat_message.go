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
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	sessionsSendChatMessageData string
	sessionsSendChatMessageInteractive bool
)

var sessionsSendChatMessageCmd = &cobra.Command{
	Use:   "send-chat-message <session-id>",
	Short: "Send Chat Message",
	Long:  `Send Chat Message

Arguments:
  session-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel sessions send-chat-message <session-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		sessionId := args[0]

		var body sdk.ChatMessageRequest

		if sessionsSendChatMessageData != "" {
			if err := json.Unmarshal([]byte(sessionsSendChatMessageData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if sessionsSendChatMessageInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Sessions.SendChatMessage(ctx, sessionId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	sessionsCmd.AddCommand(sessionsSendChatMessageCmd)

	sessionsSendChatMessageCmd.Flags().StringVarP(&sessionsSendChatMessageData, "data", "d", "", "JSON data for the request body")
	sessionsSendChatMessageCmd.Flags().BoolVarP(&sessionsSendChatMessageInteractive, "interactive", "i", false, "use interactive form input")
}
