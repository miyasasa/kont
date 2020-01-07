package main

import (
	"fmt"
	"kont/api"
	"kont/init/env"
	"kont/internal/scheduler"
	"runtime"
)

func main() {

	scheduler.ScheduleRemoteRepositories()
	_ = api.Router.Run(fmt.Sprintf(":%s", env.ServerPort))

	runtime.Goexit()

	//remoterepository.ListenRemoteRepositories()

}
