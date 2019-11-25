package util

import (
	"github.com/deckarep/golang-set"
	"github.com/stretchr/testify/assert"
	"kont/internal/common"
	"testing"
)

func TestGetReviewerFirstAvailableIntoSetConsistOneElement(t *testing.T) {
	reviewer1 := &common.Reviewer{User: common.User{Name: "atiba", DisplayName: "Atiba Hutchinson"}}

	reviewers := mapset.NewSet(reviewer1)

	reviewer := GetReviewerFirstAvailable(reviewers)

	assert.NotNil(t, reviewer)
	assert.NotNil(t, reviewer.User)
	assert.Equal(t, "atiba", reviewer.User.Name)
	assert.Equal(t, "Atiba Hutchinson", reviewer.User.DisplayName)
}

func TestGetReviewerFirstAvailableIntoSetConsistTwoElement(t *testing.T) {
	reviewer1 := &common.Reviewer{User: common.User{Name: "vida", DisplayName: "Domagoj Vida"}}
	reviewer2 := &common.Reviewer{User: common.User{Name: "atiba", DisplayName: "Atiba Hutchinson"}}

	reviewers := mapset.NewSet()
	reviewers.Add(reviewer1)
	reviewers.Add(reviewer2)

	reviewer := GetReviewerFirstAvailable(reviewers)

	assert.NotNil(t, reviewer)
	assert.NotNil(t, reviewer.User)
	assert.Equal(t, "vida", reviewer.User.Name)
	assert.Equal(t, "Domagoj Vida", reviewer.User.DisplayName)
}

func TestGetReviewerFirstAvailableWithEmptySet(t *testing.T) {
	reviewers := mapset.NewSet()

	reviewer := GetReviewerFirstAvailable(reviewers)

	assert.Nil(t, reviewer)
}

func TestGetReviewerRandomly(t *testing.T) {
	reviewer1 := &common.Reviewer{User: common.User{Name: "atiba", DisplayName: "Atiba Hutchinson"}}

	reviewers := mapset.NewSet(reviewer1)

	for i := 0; i < 10; i++ {
		reviewer := GetReviewerRandomly(reviewers)

		assert.NotNil(t, reviewer)
		assert.NotNil(t, reviewer.User)
		assert.Equal(t, "atiba", reviewer.User.Name)
		assert.Equal(t, "Atiba Hutchinson", reviewer.User.DisplayName)
	}
}

func TestGetReviewerRandomlyIntoTwoReviewers(t *testing.T) {
	reviewer1 := &common.Reviewer{User: common.User{Name: "atiba", DisplayName: "Atiba Hutchinson"}}
	reviewer2 := &common.Reviewer{User: common.User{Name: "vida", DisplayName: "Domagoj Vida"}}

	reviewers := mapset.NewSet()
	reviewers.Add(reviewer1)
	reviewers.Add(reviewer2)

	for i := 0; i < 10; i++ {
		reviewer := GetReviewerRandomly(reviewers)

		assert.NotNil(t, reviewer)
		assert.NotNil(t, reviewer.User)
		assert.NotEmpty(t, reviewer.User.Name)
		assert.NotEmpty(t, reviewer.User.DisplayName)
	}
}

func TestGetReviewerRandomlyWithEmptySet(t *testing.T) {
	reviewers := mapset.NewSet()

	reviewer := GetReviewerRandomly(reviewers)

	assert.Nil(t, reviewer)
}
