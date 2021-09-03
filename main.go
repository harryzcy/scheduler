package main

import (
	"github.com/harryzcy/scheduler/cmd"
	"github.com/harryzcy/scheduler/internal/schedulerd"
	"github.com/sevlyar/go-daemon"
)

func main() {

	if daemon.WasReborn() {
		// child
		schedulerd.Start()
	} else {
		// parent
		cmd.Execute()
	}
}
