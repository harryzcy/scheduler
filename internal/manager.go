package internal

import (
	"encoding/json"
	"log"
	"os"
)

var (
	tasks []Task

	cacheDir      = "/usr/local/ProgramCache/scheduler"
	cacheJSONFile = cacheDir + "/tasks.json"
)

func init() {
	tasks = make([]Task, 0)

	// create cache directory if not exists
	err := os.MkdirAll(cacheDir, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
}

func LoadTasks() error {
	data, err := os.ReadFile(cacheJSONFile)
	if err != nil {
		return err
	}

	var taskFile TaskFile
	err = json.Unmarshal(data, &taskFile)
	if err != nil {
		return err
	}
	tasks = taskFile.Tasks

	return nil
}

func StoreTasks() error {
	taskFile := &TaskFile{
		Tasks: tasks,
	}

	data, err := json.Marshal(taskFile)
	if err != nil {
		return err
	}

	err = os.WriteFile(cacheJSONFile, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func AddTask(task Task) {
	tasks = append(tasks, task)
	RunAdditionalTask(task)
}
