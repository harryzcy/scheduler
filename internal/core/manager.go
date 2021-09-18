package core

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	tasks map[string]*Task

	cacheDir      string
	cacheJSONFile string
)

func init() {
	tasks = make(map[string]*Task)

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

// GetCacheDir returns the directory that stores all caches
func GetCacheDir() string {
	return cacheDir
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
	setupTasks(tasks)
	run()
}

func existsTaskName(name string) bool {
	if _, ok := tasks[name]; ok {
		return true
	}
	return false
}

func AddTask(task *Task) error {
	if existsTaskName(task.Name) {
		return fmt.Errorf("task with name %s already exists", task.Name)
	}

	tasks[task.Name] = task
	err := addTask(task)
	return err
}

// ListTasks returns all tasks that are running, the error is currently always nil
func ListTasks() (map[string]*Task, error) {
	return tasks, nil
}

func RemoveTask(name string) error {
	exist := false
	var task *Task
	for _, task = range tasks {
		if name == task.Name {
			exist = true
			break
		}
	}

	if !exist {
		return fmt.Errorf("task with name %s does not exist", name)
	}

	removeTask(task)
	delete(tasks, task.Name)

	return nil
}

func RemoveAllTasks() {
	for _, task := range tasks {
		removeTask(task)
	}
	tasks = make(map[string]*Task)
}
