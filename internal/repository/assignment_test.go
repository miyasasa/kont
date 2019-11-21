package repository

import (
	"github.com/stretchr/testify/assert"
	"kont/internal/common"
	"testing"
)

func TestStageNotNilAndNilFields(t *testing.T) {
	stage := new(Stage)

	assert.NotNil(t, stage)
	assert.Equal(t, stage.Name, "")
	assert.Equal(t, stage.Policy, "")
	assert.Nil(t, stage.Reviewers)
}

func TestStageFields(t *testing.T) {
	reviewers := make([]common.Reviewer, 0)
	stage := &Stage{Name: "TestStage", Reviewers: reviewers, Policy: FIRST}

	assert.NotNil(t, stage)
	assert.Equal(t, "TestStage", stage.Name)
	assert.NotNil(t, stage.Reviewers)
	assert.Equal(t, 0, len(stage.Reviewers))
	assert.Equal(t, stage.Policy, FIRST)
}

func TestStageGetReviewerByFirstPolicy(t *testing.T) {

	stage := &Stage{Name: "TestStage", Reviewers: getDummyReviewers(), Policy: FIRST}

	reviewer := stage.GetReviewer()

	assert.NotNil(t, stage)
	assert.NotNil(t, reviewer)
	assert.NotNil(t, reviewer.User)
	assert.Equal(t, "atiba", reviewer.User.Name)
	assert.Equal(t, "Atiba Hutchinson", reviewer.User.DisplayName)
}

func TestStage_GetReviewerByRandomPolicy(t *testing.T) {
	stage := &Stage{Name: "TestStage", Reviewers: getDummyReviewers(), Policy: RANDOM}

	reviewer := stage.GetReviewer()

	assert.NotNil(t, stage)
	assert.NotNil(t, reviewer)
	assert.NotNil(t, reviewer.User)
}

func getDummyReviewers() []common.Reviewer {
	reviewer1 := common.Reviewer{User: common.User{Name: "atiba", DisplayName: "Atiba Hutchinson"}}
	reviewer2 := common.Reviewer{User: common.User{Name: "nKoudou", DisplayName: "Kevin NKoudou"}}
	reviewer3 := common.Reviewer{User: common.User{Name: "vida", DisplayName: "Domagoj Vida"}}

	return []common.Reviewer{reviewer1, reviewer2, reviewer3}

}
