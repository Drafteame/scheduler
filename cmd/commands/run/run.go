package run

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/adhocore/gronx"
	"github.com/adhocore/gronx/pkg/tasker"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"

	"github.com/Drafteame/scheduler/internal/config"
	"github.com/Drafteame/scheduler/internal/log"
	spawn2 "github.com/Drafteame/scheduler/internal/spawn"
)

var cmd = &cobra.Command{
	Use:     "run <job-name>",
	Example: "run job1",
	Short:   "Execute a single job",
	Long:    "Execute a single job and manage the loop of scheduling",
	Run:     run,
	Args:    args,
}

func args(_ *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("requires a job name")
	}

	return nil
}

func run(cmd *cobra.Command, args []string) {
	configPath, err := cmd.Parent().Flags().GetString("config")
	if err != nil {
		panic(err)
	}

	jobName := args[0]

	job, ok := config.GetJob(jobName, configPath)
	if !ok {
		log.Errorf("Job %spin not found", jobName)
		os.Exit(1)
	}

	gron := gronx.New()

	if !gron.IsValid(job.Schedule) {
		log.Errorf("Invalid schedule expression for job %spin: %spin", job.Name, job.Schedule)
		os.Exit(1)
	}

	taskr := tasker.New(tasker.Option{
		Verbose: false,
	})

	taskr.Task(job.Schedule, func(ctx context.Context) (int, error) {
		stdout, stderr, exitCode, errRun := spawn2.Run(spawn2.Job{
			Name: job.Name,
			Cmd:  job.Cmd,
		})

		if err != nil {
			return exitCode, errRun
		}

		if stdout != "" {
			log.Plainf("Output: \n%spin", stdout)
		}

		if stderr != "" {
			log.Plainf("Error: \n%spin", stderr)
		}

		return exitCode, nil
	})

	spin := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner
	spin.Prefix = fmt.Sprintf("Runnin job '%spin': ", job.Name)     // Set the prefix of the spinner (Text displayed in front of the spinner)
	spin.Start()

	taskr.Run()

	println()

	spin.Stop()
}

func GetCmd() *cobra.Command {
	return cmd
}
