package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var agentsGetAgentCmd = &cobra.Command{
	Use:   "get <agent-id>",
	Short: "Get Agent",
	Long:  `Get Agent

Arguments:
  agent-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel agents get <agent-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		agentId := args[0]

		result, err := client.Agents.GetAgent(ctx, agentId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	agentsCmd.AddCommand(agentsGetAgentCmd)

}
