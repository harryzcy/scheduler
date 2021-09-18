package cmd

import (
	"fmt"

	"github.com/harryzcy/scheduler/internal/messenger"
	"github.com/spf13/cobra"
)

var (
	name     string
	command  string
	schedule string
	once     bool
)

func init() {
	taskCmd.AddCommand(taskAddCmd)

	taskAddCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the task")
	taskAddCmd.Flags().StringVarP(&command, "command", "c", "", "Command to run")
	taskAddCmd.Flags().StringVarP(&schedule, "schedule", "s", "", "Schedule of the command")
	taskAddCmd.Flags().BoolVarP(&once, "once", "o", false, "Run the command only once")
}

var taskAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			fmt.Println("Error: missing flag: --name, -n")
			cmd.Usage()
			return
		}
		if command == "" {
			fmt.Println("Error: missing flag: --command, -c")
			cmd.Usage()
			return
		}
		if schedule == "" {
			fmt.Println("Error: missing flag: --schedule, -s")
			cmd.Usage()
			return
		}

		messenger.AddTask(name, command, schedule, once)
	},
}
