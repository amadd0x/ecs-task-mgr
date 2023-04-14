package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ecs-task-mgr",
	Short: "Deregister and delete old task definition revisions",
	Long: `ecs-task-mgr -d <family:revision>
	When executed, this cli will deregister and delete the given task definition
	and revision, then decrement the revision and continue to deregister and
	delete all previous revisions are removed.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
