package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var (
	downloadsProcessData string
	downloadsProcessInteractive bool
)

var downloadsProcessCmd = &cobra.Command{
	Use:   "process <datasource-id>",
	Short: "Process Download (sync)",
	Long:  `Process Download (sync)

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources downloads process <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body interface{}

		if downloadsProcessData != "" {
			if err := json.Unmarshal([]byte(downloadsProcessData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else {
			return fmt.Errorf("--data flag required (interactive form not available for this type)")
		}

		result, err := client.Datasources.Downloads.Process(ctx, datasourceId, &body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	downloadsCmd.AddCommand(downloadsProcessCmd)

	downloadsProcessCmd.Flags().StringVarP(&downloadsProcessData, "data", "d", "", "JSON data for the request body")
	downloadsProcessCmd.Flags().BoolVarP(&downloadsProcessInteractive, "interactive", "i", false, "use interactive form input")
}
