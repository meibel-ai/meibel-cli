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
	confidenceScoringListScoringJobsAgentExecutionId string
	confidenceScoringListScoringJobsAgentWorkflowName string
	confidenceScoringListScoringJobsAgentWorkflowVersion string
	confidenceScoringListScoringJobsAgentWorkflowExecutionId string
	confidenceScoringListScoringJobsToolId string
	confidenceScoringListScoringJobsToolInstanceId string
	confidenceScoringListScoringJobsToolExecutionId string
)

var confidenceScoringListScoringJobsCmd = &cobra.Command{
	Use:   "list-jobs",
	Short: "List Scoring Jobs",
	Long:  `List Scoring Jobs`,
	Example: "meibel confidence-scoring list-jobs --agent-name=<value> --agent-version=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := &sdk.ListScoringJobsOptions{}
		if confidenceScoringListScoringJobsAgentName != "" {
			opts.AgentName = &confidenceScoringListScoringJobsAgentName
		}
		if confidenceScoringListScoringJobsAgentVersion != "" {
			opts.AgentVersion = &confidenceScoringListScoringJobsAgentVersion
		}
		if confidenceScoringListScoringJobsAgentExecutionId != "" {
			opts.AgentExecutionId = &confidenceScoringListScoringJobsAgentExecutionId
		}
		if confidenceScoringListScoringJobsAgentWorkflowName != "" {
			opts.AgentWorkflowName = &confidenceScoringListScoringJobsAgentWorkflowName
		}
		if confidenceScoringListScoringJobsAgentWorkflowVersion != "" {
			opts.AgentWorkflowVersion = &confidenceScoringListScoringJobsAgentWorkflowVersion
		}
		if confidenceScoringListScoringJobsAgentWorkflowExecutionId != "" {
			opts.AgentWorkflowExecutionId = &confidenceScoringListScoringJobsAgentWorkflowExecutionId
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

	confidenceScoringListScoringJobsCmd.Flags().StringVarP(&confidenceScoringListScoringJobsAgentName, "agent-name", "", "", "The agent-name parameter")
	confidenceScoringListScoringJobsCmd.Flags().StringVarP(&confidenceScoringListScoringJobsAgentVersion, "agent-version", "", "", "The agent-version parameter")
	confidenceScoringListScoringJobsCmd.Flags().StringVarP(&confidenceScoringListScoringJobsAgentExecutionId, "agent-execution-id", "", "", "The agent-execution-id parameter")
	confidenceScoringListScoringJobsCmd.Flags().StringVarP(&confidenceScoringListScoringJobsAgentWorkflowName, "agent-workflow-name", "", "", "The agent-workflow-name parameter")
	confidenceScoringListScoringJobsCmd.Flags().StringVarP(&confidenceScoringListScoringJobsAgentWorkflowVersion, "agent-workflow-version", "", "", "The agent-workflow-version parameter")
	confidenceScoringListScoringJobsCmd.Flags().StringVarP(&confidenceScoringListScoringJobsAgentWorkflowExecutionId, "agent-workflow-execution-id", "", "", "The agent-workflow-execution-id parameter")
	confidenceScoringListScoringJobsCmd.Flags().StringVarP(&confidenceScoringListScoringJobsToolId, "tool-id", "", "", "The tool-id parameter")
	confidenceScoringListScoringJobsCmd.Flags().StringVarP(&confidenceScoringListScoringJobsToolInstanceId, "tool-instance-id", "", "", "The tool-instance-id parameter")
	confidenceScoringListScoringJobsCmd.Flags().StringVarP(&confidenceScoringListScoringJobsToolExecutionId, "tool-execution-id", "", "", "The tool-execution-id parameter")
}
