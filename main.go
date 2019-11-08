package main

import (
	"miya/internal/scheduler"
	"runtime"
)

func main() {

	scheduler.ListenRepositories()

	runtime.Goexit()
}
