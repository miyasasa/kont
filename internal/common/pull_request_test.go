package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPullRequest(t *testing.T) {

	pr := &PullRequest{}

	assert.NotNil(t, pr)
	assert.Empty(t, pr.Id)
	assert.Empty(t, pr.Version)
	assert.Empty(t, pr.Title)
	assert.Empty(t, pr.Description)
	assert.Empty(t, pr.Reviewers)
	assert.Empty(t, pr.Author)
	assert.Empty(t, pr.ToRef)
}

func TestReviewer(t *testing.T) {
	rv := &Reviewer{}

	assert.NotNil(t, rv)
	assert.Empty(t, rv.Priority)
	assert.Empty(t, rv.User)
	assert.Empty(t, rv.Approved)
}

func TestUser(t *testing.T) {
	usr := new(User)

	assert.NotNil(t, usr)
	assert.Empty(t, usr.Name)
	assert.Empty(t, usr.DisplayName)
}

func TestAuthor(t *testing.T) {
	author := new(Author)

	assert.NotNil(t, author)
	assert.Empty(t, author.User)
}

func TestToRef(t *testing.T) {
	toRef := new(ToRef)

	assert.NotNil(t, toRef)
	assert.Empty(t, toRef.DisplayId)
}

func TestAuthor_GetAuthorAsReviewer_GivenEmptyAuthor_ExpectEmptyReviewer(t *testing.T) {
	author := &Author{}

	rv := author.GetAuthorAsReviewer()

	assert.NotNil(t, rv)
	assert.Empty(t, rv.Priority)
	assert.Empty(t, rv.User)
	assert.Empty(t, rv.Approved)
}

func TestAuthor_GetAuthorAsReviewer_GivenAuthorWithAUser_ExpectAReviewerWithTheSameUserInfo(t *testing.T) {

	usr := User{Name: "atiba", DisplayName: "Atiba Hutchinson"}
	author := &Author{User: usr}

	rv := author.GetAuthorAsReviewer()

	assert.NotNil(t, rv)
	assert.False(t, rv.Approved)
	assert.Equal(t, 0, rv.Priority)
	assert.NotEmpty(t, rv.User)
	assert.Equal(t, usr, rv.User)
}

func TestPullRequest_IsAssignedAnyReviewer_GivenEmptyPR_ExpectFalse(t *testing.T) {
	pr := new(PullRequest)

	assert.NotNil(t, pr)
	assert.False(t, pr.IsAssignedAnyReviewer())
}

func TestPullRequest_IsAssignedAnyReviewer_GivenPRWith1RV_ExpectTrue(t *testing.T) {

	rv := &Reviewer{User: User{Name: "atiba", DisplayName: "Atiba Hutchinson"}}

	pr := new(PullRequest)
	pr.Reviewers = []*Reviewer{rv}

	assert.NotNil(t, pr)
	assert.True(t, pr.IsAssignedAnyReviewer())
}

func TestPullRequest_IsAssignedAnyReviewer_GivenPRWith2RV_ExpectTrue(t *testing.T) {

	rv1 := &Reviewer{User: User{Name: "atiba", DisplayName: "Atiba Hutchinson"}}
	rv2 := &Reviewer{User: User{Name: "cenk", DisplayName: "Cenk Tosun"}}

	pr := new(PullRequest)
	pr.Reviewers = []*Reviewer{rv1, rv2}

	assert.NotNil(t, pr)
	assert.True(t, pr.IsAssignedAnyReviewer())
}
