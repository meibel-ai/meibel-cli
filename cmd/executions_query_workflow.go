package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var (
	executionsQueryWorkflowData string
	executionsQueryWorkflowInteractive bool
)

var executionsQueryWorkflowCmd = &cobra.Command{
	Use:   "query-workflow <blueprint-instance-id> <query-name>",
	Short: "Query Workflow",
	Long:  `Query Workflow

Arguments:
  blueprint-instance-id: Unique identifier for the workflow instance
  query-name: Name of the query to execute`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel blueprints executions query-workflow <blueprint-instance-id> <query-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]
		queryName := args[1]

		var body interface{}

		if executionsQueryWorkflowData != "" {
			if err := json.Unmarshal([]byte(executionsQueryWorkflowData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else {
			return fmt.Errorf("--data flag required (interactive form not available for this type)")
		}

		result, err := client.Blueprints.Executions.QueryWorkflow(ctx, blueprintInstanceId, queryName, &body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	executionsCmd.AddCommand(executionsQueryWorkflowCmd)

	executionsQueryWorkflowCmd.Flags().StringVarP(&executionsQueryWorkflowData, "data", "d", "", "JSON data for the request body")
	executionsQueryWorkflowCmd.Flags().BoolVarP(&executionsQueryWorkflowInteractive, "interactive", "i", false, "use interactive form input")
}
