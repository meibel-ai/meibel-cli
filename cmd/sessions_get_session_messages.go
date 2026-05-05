package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel/internal/output"
)

var sessionsGetSessionMessagesCmd = &cobra.Command{
	Use:   "get-messages <session-id>",
	Short: "Get Session Messages",
	Long:  `Get Session Messages

Arguments:
  session-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel sessions get-messages <session-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		sessionId := args[0]

		result, err := client.Sessions.GetSessionMessages(ctx, sessionId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	sessionsCmd.AddCommand(sessionsGetSessionMessagesCmd)

}
