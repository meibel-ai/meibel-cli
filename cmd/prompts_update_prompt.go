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
	promptsUpdatePromptData string
	promptsUpdatePromptInteractive bool
)

var promptsUpdatePromptCmd = &cobra.Command{
	Use:   "update <prompt-id>",
	Short: "Update Prompt",
	Long:  `Update Prompt

Arguments:
  prompt-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel prompts update <prompt-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		promptId := args[0]

		var body sdk.UpdateAgentPromptRequest

		if promptsUpdatePromptData != "" {
			if err := json.Unmarshal([]byte(promptsUpdatePromptData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if promptsUpdatePromptInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.Prompts.UpdatePrompt(ctx, promptId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	promptsCmd.AddCommand(promptsUpdatePromptCmd)

	promptsUpdatePromptCmd.Flags().StringVarP(&promptsUpdatePromptData, "data", "d", "", "JSON data for the request body")
	promptsUpdatePromptCmd.Flags().BoolVarP(&promptsUpdatePromptInteractive, "interactive", "i", false, "use interactive form input")
}
