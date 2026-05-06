package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var blueprintsGetBlueprintTasksCmd = &cobra.Command{
	Use:   "get-tasks <blueprint-id>",
	Short: "Get Blueprint Tasks",
	Long:  `Get Blueprint Tasks

Arguments:
  blueprint-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints get-tasks <blueprint-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintId := args[0]

		result, err := client.Blueprints.GetBlueprintTasks(ctx, blueprintId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsCmd.AddCommand(blueprintsGetBlueprintTasksCmd)

}
