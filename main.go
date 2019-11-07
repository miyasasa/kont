package main

import (
	"miya/internal/scheduler"
	"runtime"
)

func main() {

	scheduler.Schedule()

	runtime.Goexit()
}
