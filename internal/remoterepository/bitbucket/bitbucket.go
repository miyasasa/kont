package bitbucket

import (
	"fmt"
	"log"
	"miya/internal/assignment"
	"miya/internal/common"
	"miya/internal/repository"
)

func Listen(repo *repository.Repository) {
	fmt.Println("Bitbucket-PR is listening....")
	fetchPRs(repo)

	filterToGetLatestPullRequests(repo)
	log.Printf("LatestPRCount: %d", len(repo.PR))

	assignment.Assign(repo)

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
