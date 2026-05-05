package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	promptsDeletePromptForce bool
)

var promptsDeletePromptCmd = &cobra.Command{
	Use:   "delete <prompt-id>",
	Short: "Delete Prompt",
	Long:  `Delete Prompt

Arguments:
  prompt-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel prompts delete <prompt-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		promptId := args[0]

		if !promptsDeletePromptForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		err := client.Prompts.DeletePrompt(ctx, promptId)
		if err != nil {
			return err
		}

		fmt.Println("Success")
		return nil
	},
}

func init() {
	promptsCmd.AddCommand(promptsDeletePromptCmd)

	promptsDeletePromptCmd.Flags().BoolVarP(&promptsDeletePromptForce, "force", "f", false, "skip confirmation prompt")
}
