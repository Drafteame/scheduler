# scheduler

Cron like application to schedule process that runs in background made easy

## Installation

```bash
go install github.com/Drafteame/scheduler@latest
```

## Usage

To create a new task, you need to create a new file `.scheduler.yml` in the home directory. The file should contain
a list of named jobs that you can execute in background:

```yaml
jobs:
  - name: TestJob
    schedule: "*/5 * * * *" # runs every 5 minutes
    command: "echo 'Hello, World!'"

  - name: TestJob2
    schedule: "*/10 * * * *" # runs every 10 minutes
    command: "echo 'Hello, World! 2'"
```

Also, you can specify the path to the configuration file using the `--config` flag:

```bash
scheduler --config /path/to/config.yml

# short
scheduler -c /path/to/config.yml
```

To start all jobs, run the following command:

```bash
scheduler start
```

To stop all jobs, run the following command:

```bash
scheduler stop
```

To list all jobs, run the following command:

```bash
scheduler list
```

Start a specific job and attach to its logs:

```bash
scheduler run <job-name>
```

Execute a job once:

```bash
scheduler exec <job-name>
```

Start a specific job in background:

```bash
scheduler start --job-name <job-name>
```

Stop a specific job:

```bash
scheduler stop --job-name <job-name>
```

## Development

To run the source code, you need to clone the repository and to have a full dev experiences you need to have Nix
installed in your machine.

```bash
git clone https://gihub.com/Drafteames/scheduler
cd scheduler
nix develop
```
