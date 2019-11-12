package bitbucket

import (
	"log"
	"miya/internal/common"
	"miya/internal/repository"
)

func Listen(repo *repository.Repository) {
	log.Println("Bitbucket-PR is listening....")
	fetchPRs(repo)

	filterToGetLatestPullRequests(repo)
	log.Printf("LatestPRCount: %d", len(repo.PR))

	repo.Assign()

	updatePRs(repo)
}

// An array of pull requests have not have any reviewer
func filterToGetLatestPullRequests(repo *repository.Repository) {
	prs := make([]common.PullRequest, 0)

	for _, v := range repo.PR {
		if !v.DoesHaveAnyReviewer() {
			prs = append(prs, v)
		}
	}

	repo.PR = prs
}
