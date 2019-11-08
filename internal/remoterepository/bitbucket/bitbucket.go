package bitbucket

import (
	"fmt"
	"log"
	"miya/internal/assignment"
	"miya/internal/common"
	"miya/internal/repository"
)

func Listen(repo repository.Repository) {
	fmt.Println("Bitbucket-PR is listening....")
	pullRequests := fetchPRs(repo)

	lp := getLatestPullRequests(pullRequests)
	log.Printf("LatestPRCounts: %d", len(lp))

	// refactor : update repo over pointer
	rp := assignment.Assign(repo, lp)

	updatePRsForAddingReviewers(rp)

}

// An array of pull requests has not have any reviewer
func getLatestPullRequests(prList []common.PullRequest) []common.PullRequest {
	prs := make([]common.PullRequest, 0)

	for _, v := range prList {
		if !v.DoesHaveAnyReviewer() {
			prs = append(prs, v)
		}
	}

	return prs
}
