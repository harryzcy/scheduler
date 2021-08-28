package core

import (
	"log"
	"os/exec"
	"strings"

	"github.com/robfig/cron/v3"
)

var cronRunner *cron.Cron

func init() {
	cronRunner = cron.New()
}

func parseCommand(command string) (name string, args []string) {
	args = strings.Split(command, " ")
	name = args[0]
	args = args[1:]
	return name, args
}

func runCommand(command string) {
	log.Println("running command:", command)
	name, args := parseCommand(command)
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	log.Println("finish running command:", command)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	log.Printf("combined out:\n%s\n", string(out))
}

func runInitialSetup(tasks []Task) {
	for _, task := range tasks {
		cronRunner.AddFunc(task.Schedule, func() {
			runCommand(task.Command)
		})
	}
}

func runAdditionalTask(task Task) {
	cronRunner.AddFunc(task.Schedule, func() {
		runCommand(task.Command)
	})
}
