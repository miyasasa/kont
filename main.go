package main

import (
	"log"
	"miya/storage"
)

func main() {

	// scheduler.ScheduleRemoteRepositories()
	// runtime.Goexit()

	// api.InitRouter().Run()

	err := storage.Storage.Ping()

	if err == nil {
		log.Println("PONG PONG")
	}
}
