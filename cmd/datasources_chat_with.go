package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
	"golang.org/x/term"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	datasourcesChatWithData string
	datasourcesChatWithInteractive bool
)

var datasourcesChatWithCmd = &cobra.Command{
	Use:   "chat-with",
	Short: "Chat with datasources via AI (streaming)",
	Long:  `Ask a question against one or more datasources. Returns a streaming SSE response with the AI-generated answer.`,
	Example: "meibel datasources chat-with",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		var body sdk.ChatWithDatasourceRequest

		if datasourcesChatWithData != "" {
			if err := json.Unmarshal([]byte(datasourcesChatWithData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if datasourcesChatWithInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("Message").Description("User question").Value(&body.Message),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		err := client.Datasources.ChatWith(ctx, body)
		if err != nil {
			return err
		}

		fmt.Println("Success")
		return nil
	},
}

func init() {
	datasourcesCmd.AddCommand(datasourcesChatWithCmd)

	datasourcesChatWithCmd.Flags().StringVarP(&datasourcesChatWithData, "data", "d", "", "JSON data for the request body")
	datasourcesChatWithCmd.Flags().BoolVarP(&datasourcesChatWithInteractive, "interactive", "i", false, "use interactive form input")
}
