package internal

import (
	"github.com/carlescere/scheduler"
	"miya/api/bitbucket"
)

func Schedule() {
	scheduler.Every(5).Seconds().Run(bitbucket.Listen)
}
