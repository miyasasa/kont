package model

import "kont/internal/repository"

func ConvertRepositoryToRepositoryModel(repo *repository.Repository) *RepositoryModel {
	var repoModel = &RepositoryModel{}
	repoModel.Host = repo.Host
	repoModel.ProjectName = repo.ProjectName
	repoModel.Name = repo.Name
	repoModel.DevelopmentBranch = repo.DevelopmentBranch
	repoModel.Provider = repo.Provider
	repoModel.DefaultComment = repo.DefaultComment
	repoModel.Users = repo.Users
	repoModel.Stages = repo.Stages

	return repoModel
}
