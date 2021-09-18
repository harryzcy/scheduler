package schedulerd

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/harryzcy/scheduler/internal/core"
)

func getPID() (int, error) {
	filename := filepath.Join(core.GetCacheDir(), pidFileName)
	raw, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	pid, err := strconv.Atoi(string(raw))
	if err != nil {
		return 0, err
	}

	return pid, nil
}
