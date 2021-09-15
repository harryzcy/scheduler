package schedulerd

import (
	"os"
	"path/filepath"
	"strconv"
	"syscall"
)

// Stop the schedulerd server
func Stop() error {
	filename := filepath.Join(".", pidFileName)
	raw, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	pid, err := strconv.Atoi(string(raw))
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
