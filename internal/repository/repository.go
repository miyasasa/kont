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

	//newPrs := repo.filterPullRequestsHasNotReviewer()
	//log.Printf("LatestPRCount: %v", len(newPrs))

	for i, pr := range repo.PRs {
		for _, s := range repo.Stages {
			if s.Policy == RANDOMINAVALABLE {
				busyReviewers := repo.getAssignedAndDoesNotApproveReviewers()
				owner := pr.Author.GetAuthorAsReviewer()
				repo.PRs[i].Reviewers = append(repo.PRs[i].Reviewers, s.getRandomInAvailableReviewers(busyReviewers, owner))
			} else {
				repo.PRs[i].Reviewers = append(repo.PRs[i].Reviewers, s.GetReviewer())
			}
		}
	}
}

func (repo *Repository) filterPullRequestsHasNotReviewer() []common.PullRequest {
	prs := make([]common.PullRequest, 0)

	for _, v := range repo.PRs {
		if !v.IsAssignedAnyReviewer() {
			prs = append(prs, v)
		}
	}

	return prs
}

// which reviewers are busy :)
func (repo *Repository) getAssignedAndDoesNotApproveReviewers() mapset.Set {
	reviewers := mapset.NewSet()
	for _, pr := range repo.PRs {
		reviewers = reviewers.Union(pr.GetReviewersByUnApproved())
	}

	return reviewers
}
