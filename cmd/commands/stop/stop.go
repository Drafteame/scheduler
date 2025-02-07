package stop

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/Drafteame/scheduler/internal/config"
	"github.com/Drafteame/scheduler/internal/log"
	"github.com/Drafteame/scheduler/internal/spawn"
)

var cmd = &cobra.Command{
	Use:     "stop [flags]",
	Example: "stop",
	Short:   "Stop all scheduled jobs",
	Long:    "Stop all scheduled jobs by killing the spawned processes",
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

	log.Infof("Stopping all jobs")

	conf, err := config.Load(configPath)
	if err != nil {
		panic(err)
	}

	for _, job := range conf.Jobs {
		if errSpawn := spawn.Stop(job.Name); errSpawn != nil {
			panic(errSpawn)
		}
	}
}

func runSingleJob(jobName, configPath string) {
	job, ok := config.GetJob(jobName, configPath)
	if !ok {
		panic(errors.New("fob not found"))
	}

	if err := spawn.Stop(job.Name); err != nil {
		panic(err)
	}
}

func GetCmd() *cobra.Command {
	return cmd
}
