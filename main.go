package main

import (
	"github.com/Drafteame/scheduler/cmd/commands"
	"github.com/Drafteame/scheduler/cmd/commands/exec"
	"github.com/Drafteame/scheduler/cmd/commands/list"
	"github.com/Drafteame/scheduler/cmd/commands/run"
	"github.com/Drafteame/scheduler/cmd/commands/start"
	"github.com/Drafteame/scheduler/cmd/commands/stop"
)

func main() {
	cmd := commands.GetRootCmd()

	cmd.AddCommand(list.GetCmd())
	cmd.AddCommand(exec.GetCmd())
	cmd.AddCommand(run.GetCmd())
	cmd.AddCommand(start.GetCmd())
	cmd.AddCommand(stop.GetCmd())

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
