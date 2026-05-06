package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	instancesDeleteBlueprintInstanceForce bool
)

var instancesDeleteBlueprintInstanceCmd = &cobra.Command{
	Use:   "delete-blueprint <blueprint-instance-id>",
	Short: "Delete Blueprint Instance",
	Long:  `Delete Blueprint Instance

Arguments:
  blueprint-instance-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints instances delete-blueprint <blueprint-instance-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		if !instancesDeleteBlueprintInstanceForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		err := client.Blueprints.Instances.DeleteBlueprintInstance(ctx, blueprintInstanceId)
		if err != nil {
			return err
		}

		fmt.Println("Success")
		return nil
	},
}

func init() {
	instancesCmd.AddCommand(instancesDeleteBlueprintInstanceCmd)

	instancesDeleteBlueprintInstanceCmd.Flags().BoolVarP(&instancesDeleteBlueprintInstanceForce, "force", "f", false, "skip confirmation prompt")
}
