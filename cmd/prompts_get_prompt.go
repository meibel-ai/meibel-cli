package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var promptsGetPromptCmd = &cobra.Command{
	Use:   "get <prompt-id>",
	Short: "Get Prompt",
	Long:  `Get Prompt

Arguments:
  prompt-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel prompts get <prompt-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		promptId := args[0]

		result, err := client.Prompts.GetPrompt(ctx, promptId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	promptsCmd.AddCommand(promptsGetPromptCmd)

}
