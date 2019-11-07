package internal

import "miya/api/bitbucket"

func Schedule() {
	//	scheduler.Every(5).Seconds().Run(bitbucket.Listen)
	bitbucket.Listen()
}
