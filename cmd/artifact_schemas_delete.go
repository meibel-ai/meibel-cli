package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/tui"
)

var (
	artifactSchemasDeleteForce bool
)

var artifactSchemasDeleteCmd = &cobra.Command{
	Use:   "delete <artifact-id>",
	Short: "Delete Artifact Schema",
	Long:  `Delete Artifact Schema

Arguments:
  artifact-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel artifact-schemas delete <artifact-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		artifactId := args[0]

		if !artifactSchemasDeleteForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		sp := tui.StartSpinner("Delete Artifact Schema")
		err := client.ArtifactSchemas.Delete(ctx, artifactId)
		sp.Stop()
		if err != nil {
			return err
		}

		fmt.Println("Success")
		return nil
	},
}

func init() {
	artifactSchemasCmd.AddCommand(artifactSchemasDeleteCmd)

	artifactSchemasDeleteCmd.Flags().BoolVarP(&artifactSchemasDeleteForce, "force", "f", false, "skip confirmation prompt")
}
