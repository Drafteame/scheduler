package list

import (
	"github.com/spf13/cobra"

	"github.com/Drafteame/scheduler/internal/config"
	"github.com/Drafteame/scheduler/internal/log"
)

var cmd = &cobra.Command{
	Use:     "list",
	Example: "list",
	Short:   "List all available jobs",
	Long:    "List all available jobs",
	Run:     run,
}

func run(cmd *cobra.Command, args []string) {
	configPath, err := cmd.Parent().Flags().GetString("config")
	if err != nil {
		panic(err)
	}

	conf, err := config.Load(configPath)
	if err != nil {
		panic(err)
	}

	log.Infof("\nAvailable jobs: \n")

	headers := []string{"Name", "Schedule", "Command"}

	var data [][]string

	for _, job := range conf.Jobs {
		data = append(data, []string{job.Name, job.Schedule, job.Cmd})
	}

	log.Table(headers, data)
}

func GetCmd() *cobra.Command {
	return cmd
}
