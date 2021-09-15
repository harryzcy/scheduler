package schedulerd

import (
	"log"

	"github.com/harryzcy/scheduler/internal/core"
	"github.com/harryzcy/scheduler/internal/messenger"
	"github.com/sevlyar/go-daemon"
)

var (
	pidFileName = "schedulerd.pid"
	logFileName = "schedulerd.log"
)

// start initiates cron runner and starts a server
func start() {
	core.LoadTasks()
	core.StartTasks()

	messenger.StartServer()
}

// Start runs the function in a daemon process
func Start() {
	cntxt := &daemon.Context{
		PidFileName: pidFileName,
		PidFilePerm: 0644,
		LogFileName: logFileName,
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[schedulerd]"},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Print("- - - - - - - - - - - - - - -")
	log.Print("daemon started")

	start()
}
