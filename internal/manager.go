package internal

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

var (
	tasks []Task

	cacheDir      string
	cacheJSONFile string
)

func init() {
	tasks = make([]Task, 0)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	cacheDir = filepath.Join(homeDir, "ProgramCache", "scheduler")
	cacheJSONFile = filepath.Join(cacheDir, "tasks.json")

	// create cache directory if not exists
	err = os.MkdirAll(cacheDir, os.ModePerm)
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

func StartTasks() {
	runInitialSetup(tasks)
}

func AddTask(task Task) {
	tasks = append(tasks, task)
	runAdditionalTask(task)
}
