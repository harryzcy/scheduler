package schedulerd

import (
	"fmt"
	"os"
	"syscall"
)

func status() (bool, error) {
	pid, err := getPID()
	if err != nil {
		if os.IsNotExist(err) { // if pid file does not exist, then daemon is not running
			return false, nil
		}
		return false, err
	}

	// On Unix systems, FindProcess always succeeds,
	// so the error is ignored.
	proc, _ := os.FindProcess(pid)

	// If  sig  is 0, then no signal is sent, but error checking is still per‐
	// formed; this can be used to check for the existence of a process ID  or
	// process group ID.
	err = proc.Signal(syscall.Signal(0))
	if err != nil {
		return false, nil
	}

	return true, nil
}

func Status() {
	exists, err := status()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if exists {
		fmt.Println("daemon process is running")
	} else {
		fmt.Println("daemon process is not running")
	}
}
