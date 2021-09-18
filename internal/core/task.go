package core

type Task struct {
	ID       int
	Name     string
	Command  string
	Schedule string
	Once     bool
}
