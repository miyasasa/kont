package repository

import (
	"miya/internal/common"
)

const (
	STAGE1 = "STAGE1"
	STAGE2 = "STAGE2"
	STAGE3 = "STAGE3"
)

//it will initial from remote-repository not manually
func getReviewers() map[string][]common.Reviewer {
	rv := make(map[string][]common.Reviewer)
	rv[STAGE1] = []common.Reviewer{getReviewer("ataday"), getReviewer("baydogdu"), getReviewer("huseyiny"), getReviewer("eunal")}
	rv[STAGE2] = []common.Reviewer{getReviewer("veroglu")}
	rv[STAGE3] = []common.Reviewer{getReviewer("edincer")}

	return rv
}

func getReviewer(name string) common.Reviewer {
	return common.Reviewer{User: common.User{Name: name}}
}
