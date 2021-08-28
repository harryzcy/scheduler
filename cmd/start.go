package cmd

import (
	"github.com/spf13/cobra"

	"github.com/harryzcy/scheduler/internal/schedulerd"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the deamon process",
	Run: func(cmd *cobra.Command, args []string) {
		schedulerd.Start()
	},
}
