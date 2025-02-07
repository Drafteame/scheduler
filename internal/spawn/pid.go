package spawn

import (
	"fmt"

	"github.com/Drafteame/scheduler/internal/files"
)

const pidFile = "$HOME/.scheduler/job-%s.pid"

func init() {
	if err := createSchedulerFolder(); err != nil {
		panic(err)
	}
}

func pidExists(jobName string) bool {
	pid := fmt.Sprintf(pidFile, jobName)
	return files.Exists(pid)
}

func createSchedulerFolder() error {
	if files.Exists("$HOME/.scheduler") {
		return nil
	}

	if err := files.Mkdir("$HOME/.scheduler"); err != nil {
		return fmt.Errorf("error creating scheduler folder: %w", err)
	}

	return nil
}

func cretePidFile(jobName string, pid int) error {
	pidFileName := fmt.Sprintf(pidFile, jobName)

	err := files.Write(pidFileName, []byte(fmt.Sprintf("%d", pid)))
	if err != nil {
		return fmt.Errorf("error creating pid file for job %s: %w ", jobName, err)
	}

	return nil
}

func readPid(jobName string) (int, error) {
	pidFileName := fmt.Sprintf(pidFile, jobName)

	file, err := files.Open(pidFileName)
	if err != nil {
		return 0, fmt.Errorf("error opening pid file for job %s: %w", jobName, err)
	}
	defer func() {
		_ = file.Close()
	}()

	var pid int

	_, err = fmt.Fscanf(file, "%d", &pid)
	if err != nil {
		return 0, fmt.Errorf("error reading pid file for job %s: %w", jobName, err)
	}

	return pid, nil
}

func removePidFile(jobName string) error {
	pidFileName := fmt.Sprintf(pidFile, jobName)

	if err := files.Remove(pidFileName); err != nil {
		return fmt.Errorf("error removing pid file for job %s: %w", jobName, err)
	}

	return nil
}
