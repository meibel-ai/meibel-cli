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
	executionsSendChatMessageStreamBlueprintInstanceIdChatStreamPostData string
	executionsSendChatMessageStreamBlueprintInstanceIdChatStreamPostInteractive bool
)

var executionsSendChatMessageStreamBlueprintInstanceIdChatStreamPostCmd = &cobra.Command{
	Use:   "send-chat-message-stream-blueprint-instance-id-chat-stream-post <blueprint-instance-id>",
	Short: "Send a chat message and stream the response via SSE",
	Long:  `Send a chat message to a running chat agent workflow and stream the response as Server-Sent Events.

Arguments:
  blueprint-instance-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints executions send-chat-message-stream-blueprint-instance-id-chat-stream-post <blueprint-instance-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		var body sdk.ChatMessageRequest

		if executionsSendChatMessageStreamBlueprintInstanceIdChatStreamPostData != "" {
			if err := json.Unmarshal([]byte(executionsSendChatMessageStreamBlueprintInstanceIdChatStreamPostData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if executionsSendChatMessageStreamBlueprintInstanceIdChatStreamPostInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Blueprints.Executions.SendChatMessageStreamBlueprintInstanceIdChatStreamPost(ctx, blueprintInstanceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	executionsCmd.AddCommand(executionsSendChatMessageStreamBlueprintInstanceIdChatStreamPostCmd)

	executionsSendChatMessageStreamBlueprintInstanceIdChatStreamPostCmd.Flags().StringVarP(&executionsSendChatMessageStreamBlueprintInstanceIdChatStreamPostData, "data", "d", "", "JSON data for the request body")
	executionsSendChatMessageStreamBlueprintInstanceIdChatStreamPostCmd.Flags().BoolVarP(&executionsSendChatMessageStreamBlueprintInstanceIdChatStreamPostInteractive, "interactive", "i", false, "use interactive form input")
}
