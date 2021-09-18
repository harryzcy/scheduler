package core

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Command  string `json:"command"`
	Schedule string `json:"schedule"`
	Once     bool   `json:"once"`
}
