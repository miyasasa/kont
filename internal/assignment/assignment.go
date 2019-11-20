package assignment

import (
	"kont/internal/common"
	"math/rand"
	"time"
)

func GetRandomReviewer(rvList []common.Reviewer) common.Reviewer {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(rvList))
	return rvList[index]
}

func GetFirst(rvList []common.Reviewer) common.Reviewer {
	return rvList[0]
}
