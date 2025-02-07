package start

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/Drafteame/scheduler/internal/config"
	"github.com/Drafteame/scheduler/internal/log"
	"github.com/Drafteame/scheduler/internal/spawn"
)

var cmd = &cobra.Command{
	Use:     "start [flags]",
	Example: "start",
	Short:   "Starts all scheduled jobs in background",
	Long:    "Starts all scheduled jobs in background by spawning a new process for each job",
	Run:     run,
}

var jobName string

func init() {
	cmd.Flags().StringVarP(&jobName, "job-name", "j", "", "The name of the job to start")
}

func run(cmd *cobra.Command, args []string) {
	configPath, err := cmd.Parent().Flags().GetString("config")
	if err != nil {
		panic(err)
	}

	if jobName != "" {
		runSingleJob(jobName, configPath)
		return
	}

	log.Infof("Starting all jobs")

	conf, err := config.Load(configPath)
	if err != nil {
		panic(err)
	}

	for _, job := range conf.Jobs {
		if errSpawn := spawn.Start(job.Name, configPath); errSpawn != nil {
			panic(errSpawn)
		}
	}
}

func runSingleJob(jobName, configPath string) {
	job, ok := config.GetJob(jobName, configPath)
	if !ok {
		panic(errors.New("fob not found"))
	}

	if err := spawn.Start(job.Name, configPath); err != nil {
		panic(err)
	}
}

func GetCmd() *cobra.Command {
	return cmd
}
