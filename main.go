package main

import (
	"miya/internal"
	"runtime"
)

func main() {

	internal.Schedule()

	runtime.Goexit()
}
