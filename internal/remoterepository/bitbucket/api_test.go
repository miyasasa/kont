package bitbucket

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"kont/internal/client"
	"kont/internal/common"
	"kont/internal/repository"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

func Test_AddReviewersToThePR_GivenMockHttpClient_ExpectAddedReviewersToThePRSuccessfully(t *testing.T) {

	rev1 := &common.Reviewer{Priority: 1, User: common.User{Name: "atiba"}}
	rev2 := &common.Reviewer{Priority: 2, User: common.User{Name: "noah"}}
	rev3 := &common.Reviewer{Priority: 3, User: common.User{Name: "tosun"}}

	pr := common.PullRequest{Id: 1903, Reviewers: []*common.Reviewer{rev1, rev2, rev3}}

	repo := new(repository.Repository)
	repo.Name = "testRepo"
	repo.Provider = "BITBUCKET"
	repo.Host = "http://localhost:123"
	repo.Token = "testtoken123"
	repo.PRs = []common.PullRequest{pr}

	repo.Initialize()

	bApi := NewBitbucketApi(&HttpClientMock{})

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	bApi.addReviewersToPR(0, repo)

	logString := buf.String()
	t.Log(logString)

	assert.NotEmpty(t, logString)
	assert.Contains(t, logString, "{Id:1903 Version:0 Title: Description: Reviewers:[{User:{Name:atiba}} {User:{Name:noah}} {User:{Name:tosun}}]}")
}

func Test_AddDefaultCommentToThePR_GivenMockHttpClient_ExpectAddedDefaultCommentToThePRSuccessfully(t *testing.T) {

	author := common.Author{User: common.User{Name: "atiba"}}

	pr := common.PullRequest{Id: 1903, Author: author}

	repo := new(repository.Repository)
	repo.Name = "testRepo"
	repo.Provider = "BITBUCKET"
	repo.Host = "http://localhost:123"
	repo.Token = "testtoken123"
	repo.PRs = []common.PullRequest{pr}
	repo.DefaultComment = "Hello @{{name}},\n I have an offer you can't refuse..."

	repo.Initialize()

	bApi := NewBitbucketApi(&HttpClientMock{})

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	bApi.addDefaultCommentToPR(0, repo)

	logString := buf.String()
	t.Log(logString)

	assert.NotEmpty(t, logString)
	assert.Contains(t, logString, "Added default-comment to 1903, Default-Comment: {Hello @atiba,\n I have an offer you can't refuse...}")
}

func TestUpdatePrsMethod_Given2PrsToRepoWithDefaultComment_ExpectAddedReviewersAndDefaultCommentToThePRSuccess(t *testing.T) {
	rev1 := &common.Reviewer{Priority: 1, User: common.User{Name: "atiba"}}
	rev2 := &common.Reviewer{Priority: 2, User: common.User{Name: "noah"}}
	rev3 := &common.Reviewer{Priority: 3, User: common.User{Name: "tosun"}}

	author1 := common.Author{User: common.User{Name: "fabri"}}
	pr1 := common.PullRequest{Id: 1903, Author: author1, Reviewers: []*common.Reviewer{rev1, rev2, rev3}}

	rev11 := &common.Reviewer{Priority: 1, User: common.User{Name: "oguzhan"}}
	author12 := common.Author{User: common.User{Name: "pascal"}}
	pr2 := common.PullRequest{Id: 2020, Author: author12, Reviewers: []*common.Reviewer{rev11}}

	repo := new(repository.Repository)
	repo.Name = "testRepo"
	repo.Provider = "BITBUCKET"
	repo.Host = "http://localhost:123"
	repo.Token = "testtoken123"
	repo.PRs = []common.PullRequest{pr1, pr2}
	repo.DefaultComment = "Hello @{{name}},\n I have an offer you can't refuse..."

	repo.Initialize()

	bApi := NewBitbucketApi(&HttpClientMock{})

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	bApi.updatePRs(repo)

	time.Sleep(time.Duration(1) * time.Second)

	logString := buf.String()
	t.Log(logString)

	assert.NotEmpty(t, logString)
	assert.Contains(t, logString, "{Id:1903 Version:0 Title: Description: Reviewers:[{User:{Name:atiba}} {User:{Name:noah}} {User:{Name:tosun}}]}")
	assert.Contains(t, logString, "Added default-comment to 1903, Default-Comment: {Hello @fabri,\n I have an offer you can't refuse...}")

	assert.Contains(t, logString, "{Id:2020 Version:0 Title: Description: Reviewers:[{User:{Name:oguzhan}}]}")
	assert.Contains(t, logString, "Added default-comment to 2020, Default-Comment: {Hello @pascal,\n I have an offer you can't refuse...}")

}

func TestUpdatePrsMethod_Given2PrToRepoWithoutDefaultComment_ExpectAddedReviewersToThePRSuccessfully(t *testing.T) {
	rev1 := &common.Reviewer{Priority: 1, User: common.User{Name: "atiba"}}
	rev2 := &common.Reviewer{Priority: 2, User: common.User{Name: "noah"}}
	rev3 := &common.Reviewer{Priority: 3, User: common.User{Name: "tosun"}}

	author1 := common.Author{User: common.User{Name: "fabri"}}
	pr1 := common.PullRequest{Id: 1903, Author: author1, Reviewers: []*common.Reviewer{rev1, rev2, rev3}}

	rev11 := &common.Reviewer{Priority: 1, User: common.User{Name: "oguzhan"}}
	author12 := common.Author{User: common.User{Name: "pascal"}}
	pr2 := common.PullRequest{Id: 2020, Author: author12, Reviewers: []*common.Reviewer{rev11}}

	repo := new(repository.Repository)
	repo.Name = "testRepo"
	repo.Provider = "BITBUCKET"
	repo.Host = "http://localhost:123"
	repo.Token = "testtoken123"
	repo.PRs = []common.PullRequest{pr1, pr2}

	repo.Initialize()

	bApi := NewBitbucketApi(&HttpClientMock{})

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	bApi.updatePRs(repo)

	time.Sleep(time.Duration(1) * time.Second)

	logString := buf.String()
	t.Log(logString)

	assert.NotEmpty(t, logString)
	assert.Contains(t, logString, "{Id:1903 Version:0 Title: Description: Reviewers:[{User:{Name:atiba}} {User:{Name:noah}} {User:{Name:tosun}}]}")

	assert.Contains(t, logString, "{Id:2020 Version:0 Title: Description: Reviewers:[{User:{Name:oguzhan}}]}")

}

// HttpClientMock area
type HttpClientMock struct {
	dispatcher client.Dispatcher
}

func (c *HttpClientMock) Handle(req *http.Request) {
}

func (c *HttpClientMock) HandleToInterface(req *http.Request,
	i interface{}) {
}
