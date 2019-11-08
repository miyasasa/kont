package repository

import (
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
}
