package cmd

import (
	"github.com/spf13/cobra"

	"github.com/harryzcy/scheduler/internal/schedulerd"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the status of the daemon process",
	Run: func(cmd *cobra.Command, args []string) {
		schedulerd.Status()
	},
}
