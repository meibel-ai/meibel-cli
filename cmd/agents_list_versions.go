package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	agentsListVersionsPublished string
	agentsListVersionsOffset int64
	agentsListVersionsLimit string
)

var agentsListVersionsCmd = &cobra.Command{
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

		opts := &sdk.AgentsListVersionsOptions{}
		if agentsListVersionsPublished != "" {
			opts.Published = &agentsListVersionsPublished
		}
		if agentsListVersionsOffset != 0 {
			opts.Offset = &agentsListVersionsOffset
		}
		if agentsListVersionsLimit != "" {
			opts.Limit = &agentsListVersionsLimit
		}

		iter := client.Agents.ListVersions(ctx, agentId, opts)

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
	agentsCmd.AddCommand(agentsListVersionsCmd)

	agentsListVersionsCmd.Flags().StringVarP(&agentsListVersionsPublished, "published", "", "", "If true, return only published versions. If omitted, return all versions.")
	agentsListVersionsCmd.Flags().Int64VarP(&agentsListVersionsOffset, "offset", "", 0, "Number of items to skip")
	agentsListVersionsCmd.Flags().StringVarP(&agentsListVersionsLimit, "limit", "", "", "Maximum number of items to return")
}
