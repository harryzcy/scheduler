package schedulerd

import (
	"os"
	"syscall"
)

// Stop the schedulerd server
func Stop() error {
	pid, err := getPID()
	if err != nil {
		return err
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	err = proc.Signal(syscall.SIGTERM)
	if err != nil {
		return err
	}

	return nil
}
