package main

import (
	"kont/api"
	"kont/internal/scheduler"
	"runtime"
)

func main() {

	scheduler.ScheduleRemoteRepositories()
	api.Router.Run()

	runtime.Goexit()

	//remoterepository.ListenRemoteRepositories()

}
