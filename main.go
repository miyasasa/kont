package main

import (
	"miya/internal/scheduler"
	"runtime"
)

func main() {

	scheduler.ScheduleRemoteRepositories()

	runtime.Goexit()
}
