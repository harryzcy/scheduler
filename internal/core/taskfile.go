package core

type TaskFile struct {
	Tasks map[string]*Task `json:"tasks"`
}
