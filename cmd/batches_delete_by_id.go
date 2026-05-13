package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	batchesDeleteByIdForce bool
)

var batchesDeleteByIdCmd = &cobra.Command{
	Use:   "delete-by-id <definition-id>",
	Short: "Delete Batch Definition By Id",
	Long:  `Delete Batch Definition By Id

Arguments:
  definition-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel batches delete-by-id <definition-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		definitionId := args[0]

		if !batchesDeleteByIdForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		err := client.Batches.DeleteById(ctx, definitionId)
		if err != nil {
			return err
		}

		fmt.Println("Success")
		return nil
	},
}

func init() {
	batchesCmd.AddCommand(batchesDeleteByIdCmd)

	batchesDeleteByIdCmd.Flags().BoolVarP(&batchesDeleteByIdForce, "force", "f", false, "skip confirmation prompt")
}
