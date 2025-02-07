package spawn

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/Drafteame/scheduler/internal/log"
)

const ShellToUse = "bash"

// Start starts a job by creating a new process and writing the process ID to a file.
func Start(jobName, configPath string) error {
	if pidExists(jobName) {
		log.Infof("Job %s already running...", jobName)
		return nil
	}

	log.Infof("Starting job %s...", jobName)

	args := []string{"run", jobName}

	if configPath != "" {
		args = []string{"-c", configPath, "run", jobName}
	}

	// Start the job
	cmd := exec.Command("scheduler", args...)            //nolint:gosec
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true} //nolint:exhaustruct    // detach from terminal

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error starting job %s: %w", jobName, err)
	}

	// Create pid file
	if err := cretePidFile(jobName, cmd.Process.Pid); err != nil {
		return errors.Join(err, cmd.Process.Kill())
	}

	log.Infof("Job %s started with pid %d", jobName, cmd.Process.Pid)
	return nil
}

// Stop stops a job by sending a signal to the process with the process ID.
func Stop(jobName string) error {
	remove := false

	defer func() {
		if remove {
			_ = removePidFile(jobName)
		}
	}()

	if !pidExists(jobName) {
		log.Warnf("Job %s is not running...", jobName)
		return nil
	}

	pid, err := readPid(jobName)
	if err != nil {
		return nil
	}

	log.Infof("Stopping job %s with pid %d...", jobName, pid)

	// Get the process
	process, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("error finding process for job %s: %w", jobName, err)
	}

	// Send the signal
	if err := process.Signal(syscall.SIGTERM); err != nil {
		if err.Error() == "os: process already finished" {
			remove = true
			return nil
		}

		return fmt.Errorf("error stopping job %s: %w", jobName, err)
	}

	remove = true

	log.Infof("Job %s stopped", jobName)
	return nil
}

func Run(job Job) (string, string, int, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(ShellToUse, "-c", job.Cmd) //nolint:gosec
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	exitCode := cmd.ProcessState.ExitCode()

	return stdout.String(), stderr.String(), exitCode, err
}
