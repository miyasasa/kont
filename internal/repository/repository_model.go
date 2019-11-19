package repository

import (
	"miya/internal/assignment"
	"miya/internal/common"
)

const (
	GITHUB    = "GITHUB"
	BITBUCKET = "BITBUCKET"
	GITLAB    = "GITLAB"
)

type Repository struct {
	Url         string                       `json:"url"`
	Token       string                       `json:"-"`
	ProjectName string                       `json:"projectName"`
	Name        string                       `json:"name"`
	Provider    string                       `json:"provider"`
	Users       []string                     `json:"users"`
	Reviewers   map[string][]common.Reviewer `json:"reviewers"`
	PRs         []common.PullRequest         `json:"prs"`
}

func (repo *Repository) Assign() {
	for i := range repo.PRs {
		first := assignment.GetRandomReviewer(repo.Reviewers[STAGE1])
		second := assignment.GetFirst(repo.Reviewers[STAGE2])
		third := assignment.GetFirst(repo.Reviewers[STAGE3])

		repo.PRs[i].Reviewers = []common.Reviewer{first, second, third}
	}
}
