package cmd

import (
	"github.com/harryzcy/scheduler/internal/schedulerd"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the deamon process",
	Run: func(cmd *cobra.Command, args []string) {
		schedulerd.Stop()
	},
}
