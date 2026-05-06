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
	promptsCreatePromptData string
	promptsCreatePromptInteractive bool
)

var promptsCreatePromptCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Prompt",
	Long:  `Create Prompt`,
	Example: "meibel prompts create",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		var body sdk.CreateAgentPromptRequest

		if promptsCreatePromptData != "" {
			if err := json.Unmarshal([]byte(promptsCreatePromptData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if promptsCreatePromptInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("DisplayName").Description("Human-readable name of the prompt (letters, numbers, and spaces only). Converted to kebab-case internally.").Value(&body.DisplayName),
					huh.NewInput().Title("Prompt").Description("Prompt text").Value(&body.Prompt),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.Prompts.CreatePrompt(ctx, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	promptsCmd.AddCommand(promptsCreatePromptCmd)

	promptsCreatePromptCmd.Flags().StringVarP(&promptsCreatePromptData, "data", "d", "", "JSON data for the request body")
	promptsCreatePromptCmd.Flags().BoolVarP(&promptsCreatePromptInteractive, "interactive", "i", false, "use interactive form input")
}
