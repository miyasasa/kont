package repository

import (
	"github.com/deckarep/golang-set"
	"kont/init/env"
	"kont/internal/common"
)

const (
	GITHUB    = "GITHUB"
	BITBUCKET = "BITBUCKET"
	GITLAB    = "GITLAB"
)

type Repository struct {
	FetchRepoUsersUrl    string                 `json:"fetchRepoUsersUrl"`
	FetchProjectUsersUrl string                 `json:"fetchProjectUsersUrl"`
	FetchPrsUrl          string                 `json:"fetchPrsUrl"`
	Token                string                 `json:"token"`
	ProjectName          string                 `json:"projectName"`
	Name                 string                 `json:"name"`
	Provider             string                 `json:"provider"`
	Users                map[string]common.User `json:"users"`
	Stages               []Stage                `json:"stages"`
	PRs                  []common.PullRequest   `json:"prs"`
}

func (repo *Repository) Initialize() {
	// choose according provider Bitbucket
	repo.FetchRepoUsersUrl = env.BitbucketFetchRepoUsersURL(repo.ProjectName, repo.Name)
	repo.FetchProjectUsersUrl = env.BitbucketFetchProjectUsersURL(repo.ProjectName)
	repo.FetchPrsUrl = env.BitbucketFetchPrListURL(repo.ProjectName, repo.Name)
}

func (repo *Repository) AssignReviewersToPrs() {

	busyReviewers := repo.getAssignedAndDoesNotApproveReviewers()

	//newPrs := repo.filterPullRequestsHasNotReviewer()
	//log.Printf("LatestPRCount: %v", len(newPrs))

	for i, pr := range repo.PRs {
		ownerAndReviewers := mapset.NewSet(pr.Author.GetAuthorAsReviewer())
		for _, s := range repo.Stages {
			reviewer := s.GetReviewer(busyReviewers, ownerAndReviewers)
			if reviewer == nil {
				// get reviewer from next stage. if is last stage go to first
			}

			// add reviewer to owner and reviewer
			repo.PRs[i].Reviewers = append(repo.PRs[i].Reviewers, reviewer)
		}
	}
}

/*
func (repo *Repository) getReviewer(index int, busyReviewers mapset.Set, ownerAndReviewers mapset.Set) {

	stages := append(repo.Stages[index:], repo.Stages[0:index]...)
	for i, s := range stages {

	}

	for _, s := range repo.Stages {
		reviewer := s.GetReviewer(busyReviewers, ownerAndReviewers)
		if reviewer == nil {
			// get reviewer from next stage. if is last stage go to first
		}

		// add reviewer to owner and reviewer
		repo.PRs[i].Reviewers = append(repo.PRs[i].Reviewers, *reviewer)
	}

}
*/

func (repo *Repository) filterPullRequestsHasNotReviewer() {
	prs := make([]common.PullRequest, 0)

	for _, v := range repo.PRs {
		if !v.IsAssignedAnyReviewer() {
			prs = append(prs, v)
		}
	}

	repo.PRs = prs
}

// which reviewers are busy :)
func (repo *Repository) getAssignedAndDoesNotApproveReviewers() mapset.Set {
	reviewers := mapset.NewSet()
	for _, pr := range repo.PRs {
		reviewers = reviewers.Union(pr.GetReviewersByUnApproved())
	}

	return reviewers
}
