package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitbucketFetchRepoUsersURL(t *testing.T) {

	url := BitbucketFetchRepoUsersURL("http://localhost", "TEST", "test1")

	assert.NotEmpty(t, url)
	assert.Equal(t, "http://localhost/rest/api/1.0/projects/TEST/repos/test1/permissions/users", url)
}

func TestBitbucketFetchRepoUsersURL_GivenNilParameter_ExpectInvalidURL(t *testing.T) {

	url := BitbucketFetchRepoUsersURL("http://localhost", "", "test1")

	assert.NotEmpty(t, url)
	assert.Equal(t, "http://localhost/rest/api/1.0/projects//repos/test1/permissions/users", url)
}

func TestBitbucketFetchProjectUsersURL(t *testing.T) {
	url := BitbucketFetchProjectUsersURL("http://localhost", "TEST")

	assert.NotNil(t, url)
	assert.Equal(t, "http://localhost/rest/api/1.0/projects/TEST/permissions/users", url)
}

func TestBitbucketFetchPrListURL(t *testing.T) {
	url := BitbucketFetchPrListURL("http://localhost", "TEST", "test1")

	assert.NotNil(t, url)
	assert.Equal(t, "http://localhost/rest/api/1.0/projects/TEST/repos/test1/pull-requests", url)
}
