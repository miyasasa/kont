package util

import (
	"github.com/deckarep/golang-set"
	"kont/internal/common"
)

func FilterPullRequestsHasNotReviewer(prList []common.PullRequest) []common.PullRequest {
	prs := make([]common.PullRequest, 0)

	for _, v := range prList {
		if !v.IsAssignedAnyReviewer() {
			prs = append(prs, v)
		}
	}

	return prs
}

// which reviewers are busy :)
func GetAssignedAndDoesNotApproveReviewers(prList []common.PullRequest) mapset.Set {

	reviewers := mapset.NewSet()

	for _, pr := range prList {
		for _, rv := range pr.Reviewers {

			if !rv.Approved {
				reviewers.Add(rv)
			}
		}
	}

	return reviewers
}
