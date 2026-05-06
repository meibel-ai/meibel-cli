package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var blueprintsGetBlueprintCmd = &cobra.Command{
	Use:   "get <blueprint-id>",
	Short: "Get Blueprint",
	Long:  `Get Blueprint

Arguments:
  blueprint-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints get <blueprint-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintId := args[0]

		result, err := client.Blueprints.GetBlueprint(ctx, blueprintId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsCmd.AddCommand(blueprintsGetBlueprintCmd)

}
