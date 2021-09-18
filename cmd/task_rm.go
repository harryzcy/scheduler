package cmd

import (
	"fmt"

	"github.com/harryzcy/scheduler/internal/messenger"
	"github.com/spf13/cobra"
)

var (
	rmName string
)

func init() {
	taskCmd.AddCommand(taskRmCmd)

	taskRmCmd.Flags().StringVarP(&rmName, "name", "n", "", "Name of the task to be removed")
}

var taskRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a new task",
	Run: func(cmd *cobra.Command, args []string) {
		if rmName == "" {
			fmt.Println("Error: missing flag: --name, -n")
			cmd.Usage()
			return
		}

		messenger.RemoveTask(rmName)
	},
}
