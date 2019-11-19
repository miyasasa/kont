package repository

import (
	"miya/init/env"
	"miya/internal/assignment"
	"miya/internal/common"
)

const (
	GITHUB    = "GITHUB"
	BITBUCKET = "BITBUCKET"
	GITLAB    = "GITLAB"
)

type Repository struct {
	FetchRepoUsersUrl    string                       `json:"fetchRepoUsersUrl"`
	FetchProjectUsersUrl string                       `json:"fetchProjectUsersUrl"`
	FetchPrsUrl          string                       `json:"fetchPrsUrl"`
	Token                string                       `json:"-"`
	ProjectName          string                       `json:"projectName"`
	Name                 string                       `json:"name"`
	Provider             string                       `json:"provider"`
	Users                map[string]common.User       `json:"users"`
	Reviewers            map[string][]common.Reviewer `json:"reviewers"`
	PRs                  []common.PullRequest         `json:"prs"`
}

func (repo *Repository) Initialize() {
	// choose according provider Bitbucket
	repo.FetchRepoUsersUrl = env.BitbucketFetchRepoUsersURL(repo.ProjectName, repo.Name)
	repo.FetchProjectUsersUrl = env.BitbucketFetchProjectUsersURL(repo.ProjectName)
	repo.FetchPrsUrl = env.BitbucketFetchPrListURL(repo.ProjectName, repo.Name)
	repo.Token = env.BitbucketToken
}

func (repo *Repository) Assign() {
	for i := range repo.PRs {
		first := assignment.GetRandomReviewer(repo.Reviewers[STAGE1])
		second := assignment.GetFirst(repo.Reviewers[STAGE2])
		third := assignment.GetFirst(repo.Reviewers[STAGE3])

		repo.PRs[i].Reviewers = []common.Reviewer{first, second, third}
	}
}
