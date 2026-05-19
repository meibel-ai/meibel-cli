package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	confidenceScoringListScoringJobsAgentName string
	confidenceScoringListScoringJobsAgentVersion string
	confidenceScoringListScoringJobsAgentSessionId string
	confidenceScoringListScoringJobsAgentWorkflowName string
	confidenceScoringListScoringJobsAgentWorkflowVersion string
	confidenceScoringListScoringJobsAgentWorkflowSessionId string
	confidenceScoringListScoringJobsToolId string
	confidenceScoringListScoringJobsToolInstanceId string
	confidenceScoringListScoringJobsToolExecutionId string
)

var confidenceScoringListScoringJobsCmd = &cobra.Command{
	Use:   "list-jobs",
	Short: "List scoring jobs",
	Long:  `List confidence scoring jobs, optionally filtered by identity context fields. All filters are combined with AND logic.`,
	Example: "meibel confidence-scoring list-jobs --agent-name=<value> --agent-version=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := &sdk.ConfidenceScoringListScoringJobsOptions{}
		if confidenceScoringListScoringJobsAgentName != "" {
			opts.AgentName = &confidenceScoringListScoringJobsAgentName
		}
		if confidenceScoringListScoringJobsAgentVersion != "" {
			opts.AgentVersion = &confidenceScoringListScoringJobsAgentVersion
		}
		if confidenceScoringListScoringJobsAgentSessionId != "" {
			opts.AgentSessionId = &confidenceScoringListScoringJobsAgentSessionId
		}
		if confidenceScoringListScoringJobsAgentWorkflowName != "" {
			opts.AgentWorkflowName = &confidenceScoringListScoringJobsAgentWorkflowName
		}
		if confidenceScoringListScoringJobsAgentWorkflowVersion != "" {
			opts.AgentWorkflowVersion = &confidenceScoringListScoringJobsAgentWorkflowVersion
		}
		if confidenceScoringListScoringJobsAgentWorkflowSessionId != "" {
			opts.AgentWorkflowSessionId = &confidenceScoringListScoringJobsAgentWorkflowSessionId
		}
		if confidenceScoringListScoringJobsToolId != "" {
			opts.ToolId = &confidenceScoringListScoringJobsToolId
		}
		if confidenceScoringListScoringJobsToolInstanceId != "" {
			opts.ToolInstanceId = &confidenceScoringListScoringJobsToolInstanceId
		}
		if confidenceScoringListScoringJobsToolExecutionId != "" {
			opts.ToolExecutionId = &confidenceScoringListScoringJobsToolExecutionId
		}

		result, err := client.ConfidenceScoring.ListScoringJobs(ctx, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	confidenceScoringCmd.AddCommand(confidenceScoringListScoringJobsCmd)

	confidenceScoringListScoringJobsCmd.Flags().StringVarP(&confidenceScoringListScoringJobsAgentName, "agent-name", "", "", "Filter by agent name.")
	confidenceScoringListScoringJobsCmd.Flags().StringVarP(&confidenceScoringListScoringJobsAgentVersion, "agent-version", "", "", "Filter by agent version.")
	confidenceScoringListScoringJobsCmd.Flags().StringVarP(&confidenceScoringListScoringJobsAgentSessionId, "agent-session-id", "", "", "Filter by agent session ID.")
	confidenceScoringListScoringJobsCmd.Flags().StringVarP(&confidenceScoringListScoringJobsAgentWorkflowName, "agent-workflow-name", "", "", "Filter by workflow name.")
	confidenceScoringListScoringJobsCmd.Flags().StringVarP(&confidenceScoringListScoringJobsAgentWorkflowVersion, "agent-workflow-version", "", "", "Filter by workflow version.")
	confidenceScoringListScoringJobsCmd.Flags().StringVarP(&confidenceScoringListScoringJobsAgentWorkflowSessionId, "agent-workflow-session-id", "", "", "Filter by workflow session ID.")
	confidenceScoringListScoringJobsCmd.Flags().StringVarP(&confidenceScoringListScoringJobsToolId, "tool-id", "", "", "Filter by tool identifier.")
	confidenceScoringListScoringJobsCmd.Flags().StringVarP(&confidenceScoringListScoringJobsToolInstanceId, "tool-instance-id", "", "", "Filter by tool instance identifier.")
	confidenceScoringListScoringJobsCmd.Flags().StringVarP(&confidenceScoringListScoringJobsToolExecutionId, "tool-execution-id", "", "", "Filter by tool execution identifier.")
}
