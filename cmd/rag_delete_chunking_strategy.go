package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var (
	ragDeleteChunkingStrategyForce bool
)

var ragDeleteChunkingStrategyCmd = &cobra.Command{
	Use:   "delete-chunking-strategy <datasource-id>",
	Short: "Delete Chunking Strategy",
	Long:  `Delete Chunking Strategy

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources rag delete-chunking-strategy <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		if !ragDeleteChunkingStrategyForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		result, err := client.Datasources.Rag.DeleteChunkingStrategy(ctx, datasourceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	ragCmd.AddCommand(ragDeleteChunkingStrategyCmd)

	ragDeleteChunkingStrategyCmd.Flags().BoolVarP(&ragDeleteChunkingStrategyForce, "force", "f", false, "skip confirmation prompt")
}
