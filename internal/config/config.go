package config

type Job struct {
	Name     string `yaml:"name"`
	Cmd      string `yaml:"cmd"`
	Schedule string `yaml:"schedule"`
}

type Config struct {
	Jobs []Job `yaml:"jobs"`
}

func (c Config) GetJob(name string) (Job, bool) {
	for _, job := range c.Jobs {
		if job.Name == name {
			return job, true
		}
	}

	return Job{}, false
}
