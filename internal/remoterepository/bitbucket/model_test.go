package bitbucket

import (
	"github.com/stretchr/testify/assert"
	"kont/internal/common"
	"testing"
)

func TestPullRequestDefaultCommentUpdateModel(t *testing.T) {
	defaultCommentUpdateModel := &PullRequestDefaultCommentUpdateModel{}

	assert.NotNil(t, defaultCommentUpdateModel)
	assert.Empty(t, defaultCommentUpdateModel.Text)
}

func TestNewPullRequestDefaultCommentUpdateModel_ExpectReplaceAuthorNameInText(t *testing.T) {
	text := "Hello {{name}}, May you check ..."
	author := "atiba"

	defaultCommentUpdateModel := NewPullRequestDefaultCommentUpdateModel(text, author)

	assert.NotNil(t, defaultCommentUpdateModel)
	assert.Equal(t, "Hello atiba, May you check ...", defaultCommentUpdateModel.Text)
}

func TestNewPullRequestDefaultCommentUpdateModel_GivenTextWithoutNameStatement_ExpectGivenText(t *testing.T) {
	text := "Hello, May you check ..."
	author := "atiba"

	defaultCommentUpdateModel := NewPullRequestDefaultCommentUpdateModel(text, author)

	assert.NotNil(t, defaultCommentUpdateModel)
	assert.Equal(t, "Hello, May you check ...", defaultCommentUpdateModel.Text)
}

func TestPullRequestReviewersUpdateModel(t *testing.T) {
	prUpdateModel := &PullRequestReviewersUpdateModel{}

	assert.NotNil(t, prUpdateModel)
	assert.Empty(t, prUpdateModel.Id)
	assert.Empty(t, prUpdateModel.Version)
	assert.Empty(t, prUpdateModel.Title)
	assert.Empty(t, prUpdateModel.Description)
	assert.Empty(t, prUpdateModel.Reviewers)
}

func TestReviewer(t *testing.T) {
	rv := &Reviewer{}

	assert.NotNil(t, rv)
	assert.Empty(t, rv.User)
}

func TestUser(t *testing.T) {
	usr := &User{}

	assert.NotNil(t, usr)
	assert.Empty(t, usr.Name)
}

func TestMapPullRequestToUpdateModel_Given1Reviewer_ExpectUpdateModelWith1Reviewer(t *testing.T) {
	pr := common.PullRequest{}
	pr.Id = 1234
	pr.Version = 1
	pr.Title = "Test Pr Title"
	pr.Description = "Test Description"
	pr.Reviewers = []*common.Reviewer{{User: common.User{Name: "atiba", DisplayName: "Atiba Hutchinson"}}}

	updateModel := MapPullRequestToUpdateModel(pr)

	assert.NotEmpty(t, updateModel)
	assert.Equal(t, pr.Id, updateModel.Id)
	assert.Equal(t, pr.Version, updateModel.Version)
	assert.Equal(t, pr.Title, updateModel.Title)
	assert.Equal(t, pr.Description, updateModel.Description)

	assert.NotEmpty(t, updateModel.Reviewers)
	assert.Equal(t, 1, len(updateModel.Reviewers))
	assert.Equal(t, "atiba", updateModel.Reviewers[0].User.Name)
}

func TestMapPullRequestToUpdateModel_Given2Reviewer_ExpectUpdateModelWith2Reviewer(t *testing.T) {
	pr := common.PullRequest{}
	pr.Id = 1234
	pr.Version = 1
	pr.Title = "Test Pr Title"
	pr.Description = "Test Description"

	usr1 := common.User{Name: "atiba", DisplayName: "Atiba Hutchinson"}
	usr2 := common.User{Name: "fabri", DisplayName: "Fabri"}

	pr.Reviewers = []*common.Reviewer{{User: usr1}, {User: usr2}}

	updateModel := MapPullRequestToUpdateModel(pr)

	assert.NotEmpty(t, updateModel)
	assert.Equal(t, pr.Id, updateModel.Id)
	assert.Equal(t, pr.Version, updateModel.Version)
	assert.Equal(t, pr.Title, updateModel.Title)
	assert.Equal(t, pr.Description, updateModel.Description)

	assert.NotEmpty(t, updateModel.Reviewers)
	assert.Equal(t, 2, len(updateModel.Reviewers))
}
