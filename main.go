package main

import (
	"miya/api"
)

func main() {

	// scheduler.ScheduleRemoteRepositories()
	// runtime.Goexit()

	api.InitRouter().Run()
}
