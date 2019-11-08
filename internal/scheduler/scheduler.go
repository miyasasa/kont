package scheduler

import (
	"miya/internal/remoterepository/bitbucket"
	"miya/internal/repository"
)

func Schedule() {

	// get projects ,start to listen, then assignment

	//	scheduler.Every(5).Seconds().Run(bitbucket.Listen)

	repository.InitProject()

	bitbucket.Listen()

}
