package cmd

import (
	"fmt"

	"github.com/harryzcy/scheduler/internal/messenger"
	"github.com/spf13/cobra"
)

var (
	rmName string
	rmAll  bool
)

func init() {
	taskCmd.AddCommand(taskRmCmd)

	taskRmCmd.Flags().StringVarP(&rmName, "name", "n", "", "Name of the task to be removed")
	taskRmCmd.Flags().BoolVarP(&rmAll, "all", "a", false, "Whether to remove all tasks")
}

var taskRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a new task",
	Run: func(cmd *cobra.Command, args []string) {
		if rmName != "" && rmAll {
			fmt.Println("Error: conflicting flags, you should either use --name or use --all, you can not use both")
		}

		if rmAll {
			messenger.RemoveAllTasks()
			return
		}

		if rmName == "" {
			fmt.Println("Error: missing flag: --name, -n")
			cmd.Usage()
			return
		}

		messenger.RemoveTask(rmName)
	},
}
