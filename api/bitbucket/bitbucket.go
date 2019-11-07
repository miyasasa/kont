package bitbucket

import (
	"fmt"
	"log"
	"miya/api/bitbucket/model"
)

func Listen() {
	fmt.Println("Bitbucket-PR is listening....")
	pullRequests := fetchPRs()

	requests := getLatestPullRequests(pullRequests)

	log.Printf("getLatestPullRequests : %d", len(requests))
}

// An array of pull requests has not have any reviewer
func getLatestPullRequests(prList []model.PullRequest) []model.PullRequest {
	prs := make([]model.PullRequest, 0)

	for _, v := range prList {
		if !v.DoesHaveAnyReviewer() {
			prs = append(prs, v)
		}
	}

	return prs
}
