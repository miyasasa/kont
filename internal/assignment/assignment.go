package assignment

import (
	"math/rand"
	"miya/internal/common"
	"miya/internal/repository"
	"time"
)

// move this method to Repository page
func Assign(repo repository.Repository, prs []common.PullRequest) repository.Repository {

	for i := range prs {
		first := getRandomReviewer(repo.Reviewers[repository.STAGE1])
		second := getFirst(repo.Reviewers[repository.STAGE2])
		third := getFirst(repo.Reviewers[repository.STAGE3])

		prs[i].Reviewers = []common.Reviewer{first, second, third}
	}

	repo.PR = prs

	return repo
}

func getRandomReviewer(rvList []common.Reviewer) common.Reviewer {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(rvList))
	return rvList[index]
}

func getFirst(rvList []common.Reviewer) common.Reviewer {
	return rvList[0]
}
