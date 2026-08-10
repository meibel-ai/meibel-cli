package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/meibel-ai/meibel-cli/internal/tui"
)

var datasourcesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Datasources",
	Long:  `List Datasources`,
	Example: "meibel datasources list",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		sp := tui.StartSpinner("List Datasources")
		result, err := client.Datasources.List(ctx)
		sp.Stop()
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesCmd.AddCommand(datasourcesListCmd)

}
