package core

import (
	"errors"
	"fmt"
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
		log.Printf("cmd.Run() failed with %s\n", err)
		return
	}
	log.Printf("combined out:\n%s\n", string(out))
}

func setupTasks(tasks map[string]*Task) error {
	var multierror []string

	for _, task := range tasks {
		id, err := cronRunner.AddFunc(task.Schedule, func() {
			runCommand(task.Command)
		})
		task.ID = int(id)
		if err != nil {
			multierror = append(multierror, fmt.Sprintf("failed to add task `%s`: %v", task.Name, err))
		}
	}

	return errors.New(strings.Join(multierror, ", "))
}

func addTask(task *Task) error {
	id, err := cronRunner.AddFunc(task.Schedule, func() {
		runCommand(task.Command)
	})
	task.ID = int(id)
	return err
}

func removeTask(task *Task) {
	cronRunner.Remove(cron.EntryID(task.ID))
}

func run() {
	cronRunner.Start()
}
