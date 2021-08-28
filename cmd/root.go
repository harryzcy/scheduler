package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "scheduler",
	Short: "A CLI interface that manages scheduled commands",
	Long: `Scheduler is a CLI interface that manages scheduled commands. 
	It communicates with the daemon application that deals scheduling in the background.`,
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
