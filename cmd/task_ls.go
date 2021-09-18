package cmd

import (
	"github.com/harryzcy/scheduler/internal/messenger"
	"github.com/spf13/cobra"
)

func init() {
	taskCmd.AddCommand(taskLsCmd)
}

var taskLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		messenger.ListTask()
	},
}
