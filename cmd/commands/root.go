package commands

import (
	"github.com/spf13/cobra"
)

var root = &cobra.Command{ //nolint:exhaustruct
	Use:     "scheduler <command>",
	Example: "scheduler start",
	Version: Version,
	Run:     run,
}

func init() {
	root.PersistentFlags().StringP("config", "c", "$HOME/.scheduler.yaml", "config file path")
}

func run(cmd *cobra.Command, _ []string) {
	if err := cmd.Help(); err != nil {
		panic(err)
	}
}

func GetRootCmd() *cobra.Command {
	return root
}
