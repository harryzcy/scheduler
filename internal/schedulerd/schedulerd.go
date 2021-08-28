package schedulerd

import (
	"github.com/harryzcy/scheduler/internal/core"
)

func Start() {
	core.LoadTasks()
	core.StartTasks()
}
