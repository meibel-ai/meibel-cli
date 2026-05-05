package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel/internal/output"
)

var promptsListPromptsCmd = &cobra.Command{
	Use:   "list",
	Short: "List Prompts",
	Long:  `List Prompts`,
	Example: "meibel prompts list",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		result, err := client.Prompts.ListPrompts(ctx)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	promptsCmd.AddCommand(promptsListPromptsCmd)

}
