package repository

import (
	"github.com/stretchr/testify/assert"
	"kont/internal/common"
	"os"
	"testing"
)

func TestRepositoryInstance(t *testing.T) {
	repo := new(Repository)

	assert.NotNil(t, repo)
	assert.Empty(t, repo.FetchRepoUsersUrl)
	assert.Empty(t, repo.FetchProjectUsersUrl)
	assert.Empty(t, repo.FetchPrsUrl)
	assert.Empty(t, repo.Token)
	assert.Empty(t, repo.ProjectName)
	assert.Empty(t, repo.Name)
	assert.Empty(t, repo.Provider)
	assert.Empty(t, repo.Users)
	assert.Empty(t, repo.Stages)
	assert.Empty(t, repo.PRs)
}

func TestRepository_Initialize_WithoutEnvVariables(t *testing.T) {
	repo := new(Repository)
	repo.ProjectName = "BJK"
	repo.Name = "transfer"

	repo.Initialize()

	assert.NotNil(t, repo)
	assert.NotEmpty(t, repo.FetchRepoUsersUrl)
	assert.NotEmpty(t, repo.FetchProjectUsersUrl)
	assert.NotEmpty(t, repo.FetchPrsUrl)

	assert.Equal(t, "BJKtransfer", repo.FetchRepoUsersUrl)
	assert.Equal(t, "BJK", repo.FetchProjectUsersUrl)
	assert.Equal(t, "BJKtransfer", repo.FetchPrsUrl)
}

func TestRepository_Initialize(t *testing.T) {
	repo := new(Repository)
	repo.ProjectName = "BJK"
	repo.Name = "transfer"

	setMockEnvironmentVariables()

	repo.Initialize()

	assert.NotNil(t, repo)
	assert.NotEmpty(t, repo.FetchRepoUsersUrl)
	assert.NotEmpty(t, repo.FetchProjectUsersUrl)
	assert.NotEmpty(t, repo.FetchPrsUrl)

	expectFetchRepoUsersUrl := "http://localhost/rest/api/1.0/projects/" + repo.ProjectName + "/repos/" + repo.Name + "/permissions/users"
	expectFetchProjectUsersUrl := "http://localhost/rest/api/1.0/projects/" + repo.ProjectName + "/permissions/users"
	expectFetchPrsUrl := "http://localhost/rest/api/1.0/projects/" + repo.ProjectName + "/repos/" + repo.Name + "/pull-requests"

	assert.Equal(t, expectFetchRepoUsersUrl, repo.FetchRepoUsersUrl)
	assert.Equal(t, expectFetchProjectUsersUrl, repo.FetchProjectUsersUrl)
	assert.Equal(t, expectFetchPrsUrl, repo.FetchPrsUrl)
}

func TestRepository_FilterPullRequestsHasNotReviewer_WithAllPRHasReviewer_GetNoPR(t *testing.T) {
	repo := new(Repository)

	pr1 := common.PullRequest{Id: 1903, Reviewers: getDummyReviewers()}
	pr2 := common.PullRequest{Id: 116, Reviewers: getDummyReviewers()}
	pr3 := common.PullRequest{Id: 19032, Reviewers: getDummyReviewers()}

	repo.PRs = []common.PullRequest{pr1, pr2, pr3}

	repo.filterPullRequestsHasNotReviewer()

	assert.NotNil(t, repo.PRs)
	assert.True(t, len(repo.PRs) == 0)
}

func TestRepository_FilterPullRequestsHasNotReviewer_GetOnePR(t *testing.T) {
	repo := new(Repository)

	pr1 := common.PullRequest{Id: 1903, Reviewers: getDummyReviewers()}
	pr2 := common.PullRequest{Id: 19031, Reviewers: getDummyReviewers()}
	pr3 := common.PullRequest{Id: 116}

	repo.PRs = []common.PullRequest{pr1, pr2, pr3}

	repo.filterPullRequestsHasNotReviewer()

	assert.NotNil(t, repo.PRs)
	assert.True(t, len(repo.PRs) == 1)
	assert.Equal(t, repo.PRs[0], pr3)
}

func TestRepository_FilterPullRequestsHasNotReviewer_WithAllPRHasNotReviewer_Get3PR(t *testing.T) {
	repo := new(Repository)

	pr1 := common.PullRequest{Id: 1903}
	pr2 := common.PullRequest{Id: 19031}
	pr3 := common.PullRequest{Id: 116}

	repo.PRs = []common.PullRequest{pr1, pr2, pr3}

	repo.filterPullRequestsHasNotReviewer()

	assert.NotNil(t, repo.PRs)
	assert.True(t, len(repo.PRs) == 3)
}

func TestRepository_GetAssignedAndDoesNotApproveReviewers_GivenOnePRAndAllReviewersApproved_ExpectEmptyArray(t *testing.T) {
	repo := new(Repository)

	rev1 := &common.Reviewer{Order: 1, Approved: true}
	rev2 := &common.Reviewer{Order: 2, Approved: true}
	rev3 := &common.Reviewer{Order: 3, Approved: true}

	pr := common.PullRequest{Id: 1903, Reviewers: []*common.Reviewer{rev1, rev2, rev3}}

	repo.PRs = []common.PullRequest{pr}

	busyReviewers := repo.getAssignedAndDoesNotApproveReviewers()

	assert.NotNil(t, busyReviewers)
	assert.True(t, busyReviewers.Cardinality() == 0)
}

func TestRepository_GetAssignedAndDoesNotApproveReviewers_GivenOnePRAnd2ReviewersApproved_ExpectOneReviewerInArray(t *testing.T) {
	repo := new(Repository)

	rev1 := &common.Reviewer{Order: 1, Approved: true}
	rev2 := &common.Reviewer{Order: 2, Approved: false}
	rev3 := &common.Reviewer{Order: 3, Approved: true}

	pr := common.PullRequest{Id: 1903, Reviewers: []*common.Reviewer{rev1, rev2, rev3}}

	repo.PRs = []common.PullRequest{pr}

	busyReviewers := repo.getAssignedAndDoesNotApproveReviewers()

	assert.NotNil(t, busyReviewers)
	assert.True(t, busyReviewers.Cardinality() == 1)
	assert.NotNil(t, busyReviewers.ToSlice()[0].(*common.Reviewer))
	assert.Equal(t, 2, busyReviewers.ToSlice()[0].(*common.Reviewer).Order)
}

func TestRepository_GetAssignedAndDoesNotApproveReviewers_GivenTwoPRsAndAllReviewersApproved_ExpectEmptyArray(t *testing.T) {
	repo := new(Repository)

	rev1 := &common.Reviewer{Order: 1, Approved: true}
	rev2 := &common.Reviewer{Order: 2, Approved: true}
	rev3 := &common.Reviewer{Order: 3, Approved: true}
	pr1 := common.PullRequest{Id: 1903, Reviewers: []*common.Reviewer{rev1, rev2, rev3}}

	rev4 := &common.Reviewer{Order: 4, Approved: true}
	rev5 := &common.Reviewer{Order: 5, Approved: true}

	pr2 := common.PullRequest{Id: 116, Reviewers: []*common.Reviewer{rev1, rev4, rev5}}

	repo.PRs = []common.PullRequest{pr1, pr2}

	busyReviewers := repo.getAssignedAndDoesNotApproveReviewers()

	assert.NotNil(t, busyReviewers)
	assert.True(t, busyReviewers.Cardinality() == 0)
}

func TestRepository_GetAssignedAndDoesNotApproveReviewers_GivenTwoPRsAnd2ReviewersApproved_Expect3BusyReviewerInArray(t *testing.T) {
	repo := new(Repository)

	rev1 := &common.Reviewer{Order: 1, Approved: false}
	rev2 := &common.Reviewer{Order: 2, Approved: true}
	rev3 := &common.Reviewer{Order: 3, Approved: false}
	pr1 := common.PullRequest{Id: 1903, Reviewers: []*common.Reviewer{rev1, rev2, rev3}}

	rev4 := &common.Reviewer{Order: 4, Approved: true}
	rev5 := &common.Reviewer{Order: 5, Approved: false}
	pr2 := common.PullRequest{Id: 116, Reviewers: []*common.Reviewer{rev1, rev4, rev5}}

	repo.PRs = []common.PullRequest{pr1, pr2}

	busyReviewers := repo.getAssignedAndDoesNotApproveReviewers()

	assert.NotNil(t, busyReviewers)
	assert.True(t, busyReviewers.Cardinality() == 3)
	assert.True(t, busyReviewers.Contains(rev1))
	assert.True(t, busyReviewers.Contains(rev3))
	assert.True(t, busyReviewers.Contains(rev5))
}

// In case of the reviewer available in the two PR, in one of them with approved and in the other non-approved
func TestRepository_GetAssignedAndDoesNotApproveReviewers_GivenTwoPRsAnd2ReviewersApprovedAndHaveOneCommonReviewer_Expect3BusyReviewerInArray(t *testing.T) {
	repo := new(Repository)

	rev1 := common.Reviewer{Order: 1, Approved: true}
	rev2 := &common.Reviewer{Order: 2, Approved: true}
	rev3 := &common.Reviewer{Order: 3, Approved: false}
	pr1 := common.PullRequest{Id: 1903, Reviewers: []*common.Reviewer{&rev1, rev2, rev3}}

	rev4 := &common.Reviewer{Order: 4, Approved: true}
	rev5 := &common.Reviewer{Order: 5, Approved: false}
	rev6 := rev1
	rev6.Approved = false

	pr2 := common.PullRequest{Id: 116, Reviewers: []*common.Reviewer{rev4, rev5, &rev6}}

	repo.PRs = []common.PullRequest{pr1, pr2}

	busyReviewers := repo.getAssignedAndDoesNotApproveReviewers()

	assert.NotNil(t, busyReviewers)
	assert.True(t, busyReviewers.Cardinality() == 3)
	assert.True(t, busyReviewers.Contains(&rev6))
	assert.True(t, busyReviewers.Contains(rev3))
	assert.True(t, busyReviewers.Contains(rev5))
}

func setMockEnvironmentVariables() {
	_ = os.Setenv("BITBUCKET_BASE_URL", "http://localhost/rest/api/1.0")
	_ = os.Setenv("BITBUCKET_PROJECT_PATH", "/projects/")
	_ = os.Setenv("BITBUCKET_REPOSITORY_PATH", "/repos/")
	_ = os.Setenv("BITBUCKET_PR_PATH", "/pull-requests")
	_ = os.Setenv("BITBUCKET_USER_PATH", "/permissions/users")
}
