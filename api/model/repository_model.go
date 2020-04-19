package model

import (
	"kont/internal/common"
	"kont/internal/repository"
)

type RepositoryModel struct {
	Host              string                 `json:"host" binding:"required"`
	ProjectName       string                 `json:"projectName" binding:"required"`
	Name              string                 `json:"name" binding:"required"`
	DevelopmentBranch string                 `json:"developmentBranch" binding:"required"`
	Provider          string                 `json:"provider" binding:"required"`
	DefaultComment    string                 `json:"defaultComment"`
	Users             map[string]common.User `json:"users"`
	Stages            []repository.Stage     `json:"stages"`
}
