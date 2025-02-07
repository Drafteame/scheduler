package exec

import (
	"errors"
	"os"

	"github.com/spf13/cobra"

	"github.com/Drafteame/scheduler/internal/config"
	"github.com/Drafteame/scheduler/internal/log"
	spawn2 "github.com/Drafteame/scheduler/internal/spawn"
)

var cmd = &cobra.Command{
	Use:     "exec <job-name>",
	Example: "exec job1",
	Short:   "Execute a job",
	Long:    "Execute a job and print the output, and avo",
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

	cfg, err := config.Load(configPath)
	if err != nil {
		panic(err)
	}

	jobName := args[0]

	job, ok := cfg.GetJob(jobName)
	if !ok {
		log.Errorf("Job %s not found", jobName)
		os.Exit(1)
	}

	stdout, stderr, exitCode, err := spawn2.Run(spawn2.Job{
		Name: job.Name,
		Cmd:  job.Cmd,
	})

	if err != nil {
		log.Errorf("Error: %v", err)
		os.Exit(1)
	}

	if stdout != "" {
		log.Infof("Stdout:\n%s", stdout)
	}

	if stderr != "" {
		log.Errorf("Stderr: \n%s", stderr)
	}

	if exitCode != 0 {
		log.Errorf("Error: non-zero exit code")
		os.Exit(1)
	}

	log.Infof("Exit code: %d", exitCode)
	os.Exit(0)
}

func GetCmd() *cobra.Command {
	return cmd
}
