package bitbucket

import (
	"fmt"
	"log"
	"miya/internal/common"
)

func Listen() {
	fmt.Println("Bitbucket-PR is listening....")
	pullRequests := fetchPRs()

	requests := getLatestPullRequests(pullRequests)

	log.Printf("getLatestPullRequests : %d", len(requests))
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
