package repository

import (
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

func (repo *Repository) Assign() {
	for i := range repo.PRs {
		for _, s := range repo.Stages {
			repo.PRs[i].Reviewers = append(repo.PRs[i].Reviewers, s.GetReviewer())
		}
	}
}
