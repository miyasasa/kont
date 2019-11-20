package scheduler

import (
	"github.com/carlescere/scheduler"
	"kont/internal/remoterepository"
	"log"
)

func ScheduleRemoteRepositories() {
	_, e := scheduler.Every(1).Minutes().Run(remoterepository.ListenRemoteRepositories)

	if e != nil {
		log.Printf("ScheduleRemoteRepositories:: An Error accoured to scheduler for remote repositories")
	}
}
