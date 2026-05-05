package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel/internal/output"
)

var sessionsGetSessionCmd = &cobra.Command{
	Use:   "get <session-id>",
	Short: "Get Session",
	Long:  `Get Session

Arguments:
  session-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel sessions get <session-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		sessionId := args[0]

		result, err := client.Sessions.GetSession(ctx, sessionId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	sessionsCmd.AddCommand(sessionsGetSessionCmd)

}
