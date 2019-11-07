package scheduler

import "miya/internal/repo/bitbucket"

func Schedule() {

	// get projects ,start to listen, then assignment

	//	scheduler.Every(5).Seconds().Run(bitbucket.Listen)
	bitbucket.Listen()

}
