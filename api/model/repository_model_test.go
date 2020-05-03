package model

import (
	"github.com/stretchr/testify/assert"
	"kont/internal/common"
	"kont/internal/repository"
	"testing"
)

func TestRepositoryModelInstance(t *testing.T) {
	repoModel := new(RepositoryModel)

	assert.NotNil(t, repoModel)
	assert.Empty(t, repoModel.Host)
	assert.Empty(t, repoModel.ProjectName)
	assert.Empty(t, repoModel.Name)
	assert.Empty(t, repoModel.DevelopmentBranch)
	assert.Empty(t, repoModel.Provider)
	assert.Empty(t, repoModel.DefaultComment)
	assert.Empty(t, repoModel.Users)
	assert.Empty(t, repoModel.StageModel)
}

func TestStageModelInstance(t *testing.T) {
	stageModel := new(StageModel)

	assert.NotNil(t, stageModel)
	assert.Empty(t, stageModel.Name)
	assert.Empty(t, stageModel.ReviewerModel)
	assert.Empty(t, stageModel.Policy)
}

func TestReviewerModel(t *testing.T) {
	reviewerModel := &ReviewerModel{}

	assert.NotNil(t, reviewerModel)
	assert.Empty(t, reviewerModel.Priority)
	assert.Empty(t, reviewerModel.User)
}

func TestMapReviewerModel(t *testing.T) {
	reviewer := &common.Reviewer{Priority: 1, User: common.User{Name: "atiba", DisplayName: "Atiba Hutchinson"}}

	reviewerModel := MapReviewerModel(reviewer)

	assert.NotNil(t, reviewerModel)
	assert.NotEmpty(t, reviewerModel.Priority)
	assert.Equal(t, 1, reviewerModel.Priority)
	assert.Equal(t, "atiba", reviewerModel.User.Name)
	assert.Equal(t, "Atiba Hutchinson", reviewerModel.User.DisplayName)
}

func TestMapReviewerModel_GivenEmptyUser_ExpectEmptyUserInReviewerModel(t *testing.T) {
	reviewer := &common.Reviewer{Priority: 1}

	reviewerModel := MapReviewerModel(reviewer)

	assert.NotNil(t, reviewerModel)
	assert.NotEmpty(t, reviewerModel.Priority)
	assert.Equal(t, 1, reviewerModel.Priority)
	assert.Empty(t, reviewerModel.User)
}

func TestMapStageModel_GivenEmptyStage_ExpectEmptyStageModel(t *testing.T) {
	stage := repository.Stage{}

	stageModel := MapStageModel(stage)

	assert.NotNil(t, stageModel)
	assert.Empty(t, stageModel.Name)
	assert.Empty(t, stageModel.Policy)
	assert.Empty(t, stageModel.ReviewerModel)
}

func TestMapStageModel_GivenStageWithOneReviewer_ExpectStageModelWithOneReviewer(t *testing.T) {
	reviewer := &common.Reviewer{Priority: 1, User: common.User{Name: "atiba", DisplayName: "Atiba Hutchinson"}, Approved: false}
	stage := repository.Stage{Name: "test-stage", Policy: "RANDOMINAVAILABLE", Reviewers: []*common.Reviewer{reviewer}}

	stageModel := MapStageModel(stage)

	assert.NotNil(t, stageModel)
	assert.Equal(t, "test-stage", stageModel.Name)
	assert.Equal(t, "RANDOMINAVAILABLE", stageModel.Policy)
	assert.Equal(t, 1, len(stageModel.ReviewerModel))
	assert.Equal(t, 1, stageModel.ReviewerModel[0].Priority)
	assert.Equal(t, "atiba", stageModel.ReviewerModel[0].User.Name)
	assert.Equal(t, "Atiba Hutchinson", stageModel.ReviewerModel[0].User.DisplayName)
}

func TestMapStageModel_GivenStageWith3Reviewer_ExpectStageModelWith3Reviewer(t *testing.T) {
	reviewer1 := &common.Reviewer{Priority: 0, User: common.User{Name: "abus", DisplayName: "Vincent Aboubakar"}, Approved: false}
	reviewer2 := &common.Reviewer{Priority: 1, User: common.User{Name: "talisca", DisplayName: "Anderson Talisca"}, Approved: false}
	reviewer3 := &common.Reviewer{Priority: 2, User: common.User{Name: "atiba", DisplayName: "Atiba Hutchinson"}, Approved: false}

	stage := repository.Stage{Name: "test-stage", Policy: "BYPRIORITYINAVAILABLE", Reviewers: []*common.Reviewer{reviewer1, reviewer2, reviewer3}}

	stageModel := MapStageModel(stage)

	assert.NotNil(t, stageModel)
	assert.Equal(t, "test-stage", stageModel.Name)
	assert.Equal(t, "BYPRIORITYINAVAILABLE", stageModel.Policy)
	assert.Equal(t, 3, len(stageModel.ReviewerModel))
}

func TestMapRepositoryToRepositoryModel_GivenEmptyRepository_ExpectEmptyRepositoryModel(t *testing.T) {
	repo := new(repository.Repository)

	repoModel := MapRepositoryToRepositoryModel(repo)

	assert.NotNil(t, repoModel)
	assert.Empty(t, repoModel.Name)
	assert.Empty(t, repoModel.StageModel)
}

func TestMapRepositoryToRepositoryModel(t *testing.T) {
	repo := new(repository.Repository)
	repo.Host = "http://localhost:1234"
	repo.ProjectName = "TESTPROJECT"
	repo.Token = "test-token"
	repo.Name = "test-repo"
	repo.DefaultComment = "default-comment"
	repo.DevelopmentBranch = "develop"

	reviewer := &common.Reviewer{Priority: 1, User: common.User{Name: "atiba", DisplayName: "Atiba Hutchinson"}, Approved: false}
	stage := repository.Stage{Name: "test-stage", Policy: "RANDOMINAVAILABLE", Reviewers: []*common.Reviewer{reviewer}}

	repo.Stages = []repository.Stage{stage}

	repo.Initialize()

	repoModel := MapRepositoryToRepositoryModel(repo)

	assert.NotNil(t, repoModel)
	assert.Equal(t, repo.Host, repoModel.Host)
	assert.Equal(t, repo.ProjectName, repoModel.ProjectName)
	assert.Equal(t, repo.Name, repoModel.Name)
	assert.Equal(t, repo.DevelopmentBranch, repoModel.DevelopmentBranch)
	assert.Equal(t, repo.DefaultComment, repoModel.DefaultComment)
	assert.Equal(t, 1, len(repoModel.StageModel))

}
