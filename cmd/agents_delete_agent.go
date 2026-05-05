package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	agentsDeleteAgentForce bool
)

var agentsDeleteAgentCmd = &cobra.Command{
	Use:   "delete <agent-id>",
	Short: "Delete Agent",
	Long:  `Delete Agent

Arguments:
  agent-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel agents delete <agent-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		agentId := args[0]

		if !agentsDeleteAgentForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		err := client.Agents.DeleteAgent(ctx, agentId)
		if err != nil {
			return err
		}

		fmt.Println("Success")
		return nil
	},
}

func init() {
	agentsCmd.AddCommand(agentsDeleteAgentCmd)

	agentsDeleteAgentCmd.Flags().BoolVarP(&agentsDeleteAgentForce, "force", "f", false, "skip confirmation prompt")
}
