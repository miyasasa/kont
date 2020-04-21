package model

import (
	"kont/internal/common"
)

type RepositoryModel struct {
	Host              string                 `json:"host" binding:"required"`
	ProjectName       string                 `json:"projectName" binding:"required"`
	Name              string                 `json:"name" binding:"required"`
	DevelopmentBranch string                 `json:"developmentBranch" binding:"required"`
	Provider          string                 `json:"provider" binding:"required"`
	DefaultComment    string                 `json:"defaultComment"`
	Users             map[string]common.User `json:"users"`
	StageModel        []StageModel           `json:"stages"`
}

type StageModel struct {
	Name          string          `json:"name"`
	ReviewerModel []ReviewerModel `json:"reviewers"`
	Policy        string          `json:"policy"`
}

type ReviewerModel struct {
	Priority int         `json:"priority"`
	User     common.User `json:"user"`
}
