package scheduler

import (
	"github.com/carlescere/scheduler"
	"log"
	"miya/internal/remoterepository"
)

func ScheduleRemoteRepositories() {
	_, e := scheduler.Every(10).Seconds().Run(remoterepository.ListenRemoteRepositories)

	if e != nil {
		log.Printf("ScheduleRemoteRepositories:: An Error accoured to scheduler for remote repositories")
	}
}
