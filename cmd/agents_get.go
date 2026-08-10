package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/meibel-ai/meibel-cli/internal/tui"
)

var agentsGetCmd = &cobra.Command{
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

		sp := tui.StartSpinner("Get Agent")
		result, err := client.Agents.Get(ctx, agentId)
		sp.Stop()
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	agentsCmd.AddCommand(agentsGetCmd)

}
