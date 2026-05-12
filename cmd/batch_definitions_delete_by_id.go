package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	batchDefinitionsDeleteByIdForce bool
)

var batchDefinitionsDeleteByIdCmd = &cobra.Command{
	Use:   "delete-by-id <definition-id>",
	Short: "Delete Batch Definition By Id",
	Long:  `Delete Batch Definition By Id

Arguments:
  definition-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel batch-definitions delete-by-id <definition-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		definitionId := args[0]

		if !batchDefinitionsDeleteByIdForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		err := client.BatchDefinitions.DeleteById(ctx, definitionId)
		if err != nil {
			return err
		}

		fmt.Println("Success")
		return nil
	},
}

func init() {
	batchDefinitionsCmd.AddCommand(batchDefinitionsDeleteByIdCmd)

	batchDefinitionsDeleteByIdCmd.Flags().BoolVarP(&batchDefinitionsDeleteByIdForce, "force", "f", false, "skip confirmation prompt")
}
