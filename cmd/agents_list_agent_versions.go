package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	agentsListAgentVersionsPublished string
	agentsListAgentVersionsOffset int64
	agentsListAgentVersionsLimit string
)

var agentsListAgentVersionsCmd = &cobra.Command{
	Use:   "list-versions <agent-id>",
	Short: "List Agent Versions",
	Long:  `List Agent Versions

Arguments:
  agent-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel agents list-versions <agent-id> --published=<value> --offset=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		agentId := args[0]

		opts := &sdk.ListAgentVersionsOptions{}
		if agentsListAgentVersionsPublished != "" {
			opts.Published = &agentsListAgentVersionsPublished
		}
		if agentsListAgentVersionsOffset != 0 {
			opts.Offset = &agentsListAgentVersionsOffset
		}
		if agentsListAgentVersionsLimit != "" {
			opts.Limit = &agentsListAgentVersionsLimit
		}

		iter := client.Agents.ListAgentVersions(ctx, agentId, opts)

		var items []interface{}
		for iter.Next(ctx) {
			items = append(items, iter.Item())
		}
		if err := iter.Err(); err != nil {
			return err
		}

		return output.Print(items)
	},
}

func init() {
	agentsCmd.AddCommand(agentsListAgentVersionsCmd)

	agentsListAgentVersionsCmd.Flags().StringVarP(&agentsListAgentVersionsPublished, "published", "", "", "If true, return only published versions. If omitted, return all versions.")
	agentsListAgentVersionsCmd.Flags().Int64VarP(&agentsListAgentVersionsOffset, "offset", "", 0, "Number of items to skip")
	agentsListAgentVersionsCmd.Flags().StringVarP(&agentsListAgentVersionsLimit, "limit", "", "", "Maximum number of items to return")
}
