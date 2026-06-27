package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var (
	agentsSessionsCreateByNameData string
	agentsSessionsCreateByNameInteractive bool
)

var agentsSessionsCreateByNameCmd = &cobra.Command{
	Use:   "create-by-name <name>",
	Short: "Create Session By Name",
	Long:  `Start a session against the latest published version of an agent by name.

Resolves the current latest published version at runtime — callers do not
need to track a specific agent ID or version. Returns 404 if no published
version exists for the given agent name.

Arguments:
  name: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel agents agents-sessions create-by-name <name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		name := args[0]

		var body interface{}

		if agentsSessionsCreateByNameData != "" {
			if err := json.Unmarshal([]byte(agentsSessionsCreateByNameData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else {
			return fmt.Errorf("--data flag required (interactive form not available for this type)")
		}

		result, err := client.Agents.Sessions.CreateByName(ctx, name, &body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	agentsSessionsCmd.AddCommand(agentsSessionsCreateByNameCmd)

	agentsSessionsCreateByNameCmd.Flags().StringVarP(&agentsSessionsCreateByNameData, "data", "d", "", "JSON data for the request body")
	agentsSessionsCreateByNameCmd.Flags().BoolVarP(&agentsSessionsCreateByNameInteractive, "interactive", "i", false, "use interactive form input")
}
