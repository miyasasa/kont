package model

import (
	"kont/internal/common"
	"kont/internal/repository"
)

func MapRepositoryToRepositoryModel(repo *repository.Repository) *RepositoryModel {
	var repoModel = &RepositoryModel{}
	repoModel.Host = repo.Host
	repoModel.ProjectName = repo.ProjectName
	repoModel.Name = repo.Name
	repoModel.DevelopmentBranch = repo.DevelopmentBranch
	repoModel.Provider = repo.Provider
	repoModel.DefaultComment = repo.DefaultComment
	repoModel.Users = repo.Users

	stageModels := make([]StageModel, 0, len(repo.Stages))
	for _, s := range repo.Stages {
		sm := MapStageModel(s)
		stageModels = append(stageModels, sm)
	}

	repoModel.StageModel = stageModels

	return repoModel
}

func MapStageModel(stage repository.Stage) StageModel {
	var stageModel = StageModel{Name: stage.Name, Policy: stage.Policy}

	reviewerModels := make([]ReviewerModel, 0, len(stage.Reviewers))
	for _, r := range stage.Reviewers {
		rm := MapReviewerModel(r)
		reviewerModels = append(reviewerModels, rm)
	}

	stageModel.ReviewerModel = reviewerModels

	return stageModel
}

func MapReviewerModel(reviewer *common.Reviewer) ReviewerModel {
	return ReviewerModel{Priority: reviewer.Priority, User: reviewer.User}
}
