package repository

import (
	"github.com/deckarep/golang-set"
	"github.com/stretchr/testify/assert"
	"kont/internal/common"
	"testing"
)

func TestStageNotNilAndNilFields(t *testing.T) {
	stage := new(Stage)

	assert.NotNil(t, stage)
	assert.Empty(t, stage.Name)
	assert.Empty(t, stage.Policy)
	assert.Nil(t, stage.Reviewers)
}

func TestStageFields(t *testing.T) {
	reviewers := make([]*common.Reviewer, 0)
	stage := &Stage{Name: "TestStage", Reviewers: reviewers, Policy: BYORDERINAVAILABLE}

	assert.NotNil(t, stage)
	assert.Equal(t, "TestStage", stage.Name)
	assert.NotNil(t, stage.Reviewers)
	assert.Equal(t, 0, len(stage.Reviewers))
	assert.Equal(t, stage.Policy, BYORDERINAVAILABLE)
}

func TestStageGetReviewerByOrderIn2AvailableWithOneBusyAndOneReviewer(t *testing.T) {

	reviewers := getDummyReviewers()
	stage := &Stage{Name: "TestStage", Reviewers: reviewers, Policy: BYORDERINAVAILABLE}

	assert.True(t, len(reviewers) > 3)

	busyReviewers := mapset.NewSet(reviewers[1])
	ownerAndReviewers := mapset.NewSet(reviewers[2])

	reviewer := stage.GetReviewer(busyReviewers, ownerAndReviewers)

	assert.NotNil(t, stage)
	assert.NotNil(t, reviewer)
	assert.NotNil(t, reviewer.User)
	assert.Equal(t, "atiba", reviewer.User.Name)
	assert.Equal(t, "Atiba Hutchinson", reviewer.User.DisplayName)
	assert.False(t, busyReviewers.Contains(reviewer))
	assert.False(t, ownerAndReviewers.Contains(reviewer))
}

func TestStageGetReviewerByOrderInOneAvailableWith2BusyAndOneExistReviewer(t *testing.T) {

	reviewers := getDummyReviewers()
	stage := &Stage{Name: "TestStage", Reviewers: reviewers, Policy: BYORDERINAVAILABLE}

	assert.True(t, len(reviewers) > 3)

	busyReviewers := mapset.NewSet()
	busyReviewers.Add(reviewers[0])
	busyReviewers.Add(reviewers[1])

	ownerAndReviewers := mapset.NewSet(reviewers[3])

	reviewer := stage.GetReviewer(busyReviewers, ownerAndReviewers)

	assert.NotNil(t, stage)
	assert.NotNil(t, reviewer)
	assert.NotNil(t, reviewer.User)
	assert.Equal(t, "vida", reviewer.User.Name)
	assert.Equal(t, "Domagoj Vida", reviewer.User.DisplayName)
}

func TestStageGetReviewerByOrderIn1AvailableWithOneBusyAnd2ExistReviewer(t *testing.T) {

	reviewers := getDummyReviewers()
	stage := &Stage{Name: "TestStage", Reviewers: reviewers, Policy: BYORDERINAVAILABLE}

	assert.True(t, len(reviewers) > 3)

	busyReviewers := mapset.NewSet(reviewers[0])

	ownerAndReviewers := mapset.NewSet()
	ownerAndReviewers.Add(reviewers[1])
	ownerAndReviewers.Add(reviewers[2])

	reviewer := stage.GetReviewer(busyReviewers, ownerAndReviewers)

	assert.NotNil(t, stage)
	assert.NotNil(t, reviewer)
	assert.NotNil(t, reviewer.User)
	assert.Equal(t, "gokhan", reviewer.User.Name)
	assert.Equal(t, "Gökhan Gönül", reviewer.User.DisplayName)
}

func TestStageGetReviewerByOrderIn0AvailableWith2BusyAnd2ExistReviewerGetReviewerFromBusiesRandomly(t *testing.T) {

	reviewers := getDummyReviewers()
	stage := &Stage{Name: "TestStage", Reviewers: reviewers, Policy: BYORDERINAVAILABLE}

	assert.True(t, len(reviewers) > 3)

	busyReviewers := mapset.NewSet()
	busyReviewers.Add(reviewers[1])
	busyReviewers.Add(reviewers[2])

	ownerAndReviewers := mapset.NewSet()
	ownerAndReviewers.Add(reviewers[0])
	ownerAndReviewers.Add(reviewers[3])

	reviewer := stage.GetReviewer(busyReviewers, ownerAndReviewers)

	assert.NotNil(t, stage)
	assert.NotNil(t, reviewer)
	assert.NotNil(t, reviewer.User)
	assert.True(t, busyReviewers.Contains(reviewer))
	assert.False(t, ownerAndReviewers.Contains(reviewer))
}

// in case of zero available reviewer and one busy reviewer is also owner
func TestStageGetReviewerByOrderIn0AvailableWithTheSameOneBusyAndOneExistReviewerGetNil(t *testing.T) {

	reviewers := getDummyReviewers()
	stage := &Stage{Name: "TestStage", Reviewers: make([]*common.Reviewer, 0), Policy: BYORDERINAVAILABLE}

	assert.True(t, len(reviewers) > 3)

	busyReviewers := mapset.NewSet(reviewers[0])
	ownerAndReviewers := mapset.NewSet(reviewers[0])

	reviewer := stage.GetReviewer(busyReviewers, ownerAndReviewers)

	assert.NotNil(t, stage)
	assert.Nil(t, reviewer)
	assert.False(t, busyReviewers.Contains(reviewer))
	assert.False(t, ownerAndReviewers.Contains(reviewer))
}

func TestStageGetReviewerByRandomInOneAvailableWith2BusyAndOneExistReviewer(t *testing.T) {

	reviewers := getDummyReviewers()
	stage := &Stage{Name: "TestStage", Reviewers: reviewers, Policy: RANDOMINAVAILABLE}

	assert.True(t, len(reviewers) > 3)

	busyReviewers := mapset.NewSet()
	busyReviewers.Add(reviewers[0])
	busyReviewers.Add(reviewers[1])

	ownerAndReviewers := mapset.NewSet(reviewers[3])

	reviewer := stage.GetReviewer(busyReviewers, ownerAndReviewers)

	assert.NotNil(t, stage)
	assert.NotNil(t, reviewer)
	assert.NotNil(t, reviewer.User)
	assert.Equal(t, "vida", reviewer.User.Name)
	assert.Equal(t, "Domagoj Vida", reviewer.User.DisplayName)
}

func TestStageGetReviewerRandomlyIn2AvailableWithOneBusyAndOneReviewer(t *testing.T) {

	reviewers := getDummyReviewers()
	stage := &Stage{Name: "TestStage", Reviewers: reviewers, Policy: RANDOMINAVAILABLE}

	assert.True(t, len(reviewers) > 3)

	busyReviewers := mapset.NewSet(reviewers[1])
	ownerAndReviewers := mapset.NewSet(reviewers[2])

	availableReviewers := []*common.Reviewer{reviewers[0], reviewers[3]}

	reviewer := stage.GetReviewer(busyReviewers, ownerAndReviewers)

	assert.NotNil(t, stage)
	assert.NotNil(t, reviewer)
	assert.NotNil(t, reviewer.User)
	assert.Contains(t, availableReviewers, reviewer)
	assert.False(t, busyReviewers.Contains(reviewer))
	assert.False(t, ownerAndReviewers.Contains(reviewer))
}

func TestStageGetReviewerByOrderIn0StageReviewerWithOneBusyAnd3ExistReviewerGetReviewerFromBusiesRandomly(t *testing.T) {

	reviewers := getDummyReviewers()
	stage := &Stage{Name: "TestStage", Reviewers: reviewers, Policy: RANDOMINAVAILABLE}

	assert.True(t, len(reviewers) > 3)

	busyReviewers := mapset.NewSet(reviewers[1])

	ownerAndReviewers := mapset.NewSet()
	ownerAndReviewers.Add(reviewers[0])
	ownerAndReviewers.Add(reviewers[3])
	ownerAndReviewers.Add(reviewers[2])

	reviewer := stage.GetReviewer(busyReviewers, ownerAndReviewers)

	assert.NotNil(t, stage)
	assert.NotNil(t, reviewer)
	assert.NotNil(t, reviewer.User)
	assert.Equal(t, "nKoudou", reviewer.User.Name)
	assert.Equal(t, "Kevin NKoudou", reviewer.User.DisplayName)
	assert.True(t, busyReviewers.Contains(reviewer))
	assert.False(t, ownerAndReviewers.Contains(reviewer))
}

func TestStageGetReviewerByOrderIn2StageReviewer2BusiesAlsoTheSame2ExistReviewerGetNil(t *testing.T) {

	reviewers := getDummyReviewers()
	stage := &Stage{Name: "TestStage", Reviewers: []*common.Reviewer{reviewers[0], reviewers[1]}, Policy: RANDOMINAVAILABLE}

	assert.True(t, len(reviewers) > 3)

	busyReviewers := mapset.NewSet()
	busyReviewers.Add(reviewers[0])
	busyReviewers.Add(reviewers[1])

	ownerAndReviewers := mapset.NewSet()
	ownerAndReviewers.Add(reviewers[0])
	ownerAndReviewers.Add(reviewers[1])

	reviewer := stage.GetReviewer(busyReviewers, ownerAndReviewers)

	assert.NotNil(t, stage)
	assert.Nil(t, reviewer)
	assert.False(t, busyReviewers.Contains(reviewer))
	assert.False(t, ownerAndReviewers.Contains(reviewer))
}

func getDummyReviewers() []*common.Reviewer {
	reviewer1 := &common.Reviewer{User: common.User{Name: "atiba", DisplayName: "Atiba Hutchinson"}, Order: 1}
	reviewer2 := &common.Reviewer{User: common.User{Name: "nKoudou", DisplayName: "Kevin NKoudou"}, Order: 2}
	reviewer3 := &common.Reviewer{User: common.User{Name: "vida", DisplayName: "Domagoj Vida"}, Order: 3}
	reviewer4 := &common.Reviewer{User: common.User{Name: "gokhan", DisplayName: "Gökhan Gönül"}, Order: 4}

	return []*common.Reviewer{reviewer1, reviewer2, reviewer3, reviewer4}
}
