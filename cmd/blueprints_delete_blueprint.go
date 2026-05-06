package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var (
	blueprintsDeleteBlueprintForce bool
)

var blueprintsDeleteBlueprintCmd = &cobra.Command{
	Use:   "delete <blueprint-id>",
	Short: "Delete Blueprint",
	Long:  `Delete Blueprint

Arguments:
  blueprint-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints delete <blueprint-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintId := args[0]

		if !blueprintsDeleteBlueprintForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		result, err := client.Blueprints.DeleteBlueprint(ctx, blueprintId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsCmd.AddCommand(blueprintsDeleteBlueprintCmd)

	blueprintsDeleteBlueprintCmd.Flags().BoolVarP(&blueprintsDeleteBlueprintForce, "force", "f", false, "skip confirmation prompt")
}
