package repository

import (
	"miya/internal/assignment"
	"miya/internal/common"
)

const (
	GITHUB    = "Github"
	BITBUCKET = "Bitbucket"
	GITLAB    = "Gitlab"
)

type Repository struct {
	FetchPrURL  string
	Token       string
	ProjectName string
	Name        string
	Provider    string
	Users       []string
	Reviewers   map[string][]common.Reviewer
	PR          []common.PullRequest
}

func (repo *Repository) Assign() {
	for i := range repo.PR {
		first := assignment.GetRandomReviewer(repo.Reviewers[STAGE1])
		second := assignment.GetFirst(repo.Reviewers[STAGE2])
		third := assignment.GetFirst(repo.Reviewers[STAGE3])

		repo.PR[i].Reviewers = []common.Reviewer{first, second, third}
	}
}
