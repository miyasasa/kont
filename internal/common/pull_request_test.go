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
