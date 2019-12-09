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

func TestRepository_Initialize_ExpectSuccess(t *testing.T) {
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

	rev1 := &common.Reviewer{Order: 1, Approved: true, User: common.User{Name: "rev1"}}
	rev2 := &common.Reviewer{Order: 2, Approved: true, User: common.User{Name: "rev2"}}
	rev3 := &common.Reviewer{Order: 3, Approved: true, User: common.User{Name: "rev3"}}

	pr := common.PullRequest{Id: 1903, Reviewers: []*common.Reviewer{rev1, rev2, rev3}}

	repo.PRs = []common.PullRequest{pr}

	busyReviewers := repo.getAssignedAndDoesNotApproveReviewers()
	repo.Stages = []Stage{{Reviewers: []*common.Reviewer{rev1, rev2, rev3}}}

	assert.NotNil(t, busyReviewers)
	assert.True(t, busyReviewers.Cardinality() == 0)
}

func TestRepository_GetAssignedAndDoesNotApproveReviewers_GivenOnePRAnd2ReviewersApprovedAndThirdReviewerNotAvailableInTheStage_ExpectEmptyArray(t *testing.T) {
	repo := new(Repository)

	rev1 := &common.Reviewer{Order: 1, Approved: true, User: common.User{Name: "rev1"}}
	rev2 := &common.Reviewer{Order: 2, Approved: true, User: common.User{Name: "rev2"}}
	rev3 := &common.Reviewer{Order: 3, Approved: false, User: common.User{Name: "rev3"}}

	pr := common.PullRequest{Id: 1903, Reviewers: []*common.Reviewer{rev1, rev2, rev3}}

	repo.PRs = []common.PullRequest{pr}
	repo.Stages = []Stage{{Reviewers: []*common.Reviewer{rev1, rev2}}}

	busyReviewers := repo.getAssignedAndDoesNotApproveReviewers()

	assert.NotNil(t, busyReviewers)
	assert.True(t, busyReviewers.Cardinality() == 0)
}

func TestRepository_GetAssignedAndDoesNotApproveReviewers_GivenOnePRAnd2ReviewersApproved_ExpectOneReviewerInArray(t *testing.T) {
	repo := new(Repository)

	rev1 := &common.Reviewer{Order: 1, Approved: true, User: common.User{Name: "rev1"}}
	rev2 := &common.Reviewer{Order: 2, Approved: false, User: common.User{Name: "rev2"}}
	rev3 := &common.Reviewer{Order: 3, Approved: true, User: common.User{Name: "rev3"}}

	pr := common.PullRequest{Id: 1903, Reviewers: []*common.Reviewer{rev1, rev2, rev3}}

	repo.PRs = []common.PullRequest{pr}
	repo.Stages = []Stage{{Reviewers: []*common.Reviewer{rev1, rev2, rev3}}}

	busyReviewers := repo.getAssignedAndDoesNotApproveReviewers()

	assert.NotNil(t, busyReviewers)
	assert.True(t, busyReviewers.Cardinality() == 1)
	assert.NotNil(t, busyReviewers.ToSlice()[0].(*common.Reviewer))
	assert.Equal(t, 2, busyReviewers.ToSlice()[0].(*common.Reviewer).Order)
}

func TestRepository_GetAssignedAndDoesNotApproveReviewers_GivenTwoPRsAndAllReviewersApproved_ExpectEmptyArray(t *testing.T) {
	repo := new(Repository)

	rev1 := &common.Reviewer{Order: 1, Approved: true, User: common.User{Name: "rev1"}}
	rev2 := &common.Reviewer{Order: 2, Approved: true, User: common.User{Name: "rev2"}}
	rev3 := &common.Reviewer{Order: 3, Approved: true, User: common.User{Name: "rev3"}}
	pr1 := common.PullRequest{Id: 1903, Reviewers: []*common.Reviewer{rev1, rev2, rev3}}

	rev4 := &common.Reviewer{Order: 4, Approved: true, User: common.User{Name: "rev4"}}
	rev5 := &common.Reviewer{Order: 5, Approved: true, User: common.User{Name: "rev5"}}

	pr2 := common.PullRequest{Id: 116, Reviewers: []*common.Reviewer{rev1, rev4, rev5}}

	repo.PRs = []common.PullRequest{pr1, pr2}
	repo.Stages = []Stage{{Reviewers: []*common.Reviewer{rev1, rev2, rev4}}, {Reviewers: []*common.Reviewer{rev5, rev3}}}

	busyReviewers := repo.getAssignedAndDoesNotApproveReviewers()

	assert.NotNil(t, busyReviewers)
	assert.True(t, busyReviewers.Cardinality() == 0)
}

func TestRepository_GetAssignedAndDoesNotApproveReviewers_GivenTwoPRsAnd2ReviewersApproved_Expect3BusyReviewerInArray(t *testing.T) {
	repo := new(Repository)

	rev1 := &common.Reviewer{Order: 1, Approved: false, User: common.User{Name: "rev1"}}
	rev2 := &common.Reviewer{Order: 2, Approved: true, User: common.User{Name: "rev2"}}
	rev3 := &common.Reviewer{Order: 3, Approved: false, User: common.User{Name: "rev3"}}
	pr1 := common.PullRequest{Id: 1903, Reviewers: []*common.Reviewer{rev1, rev2, rev3}}

	rev4 := &common.Reviewer{Order: 4, Approved: true, User: common.User{Name: "rev4"}}
	rev5 := &common.Reviewer{Order: 5, Approved: false, User: common.User{Name: "rev5"}}
	pr2 := common.PullRequest{Id: 116, Reviewers: []*common.Reviewer{rev1, rev4, rev5}}

	repo.PRs = []common.PullRequest{pr1, pr2}
	repo.Stages = []Stage{{Reviewers: []*common.Reviewer{rev2, rev4, rev5}}, {Reviewers: []*common.Reviewer{rev1, rev3}}}

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

	rev1 := common.Reviewer{Order: 1, Approved: true, User: common.User{Name: "rev1"}}
	rev2 := &common.Reviewer{Order: 2, Approved: true, User: common.User{Name: "rev2"}}
	rev3 := &common.Reviewer{Order: 3, Approved: false, User: common.User{Name: "rev3"}}
	pr1 := common.PullRequest{Id: 1903, Reviewers: []*common.Reviewer{&rev1, rev2, rev3}}

	rev4 := &common.Reviewer{Order: 4, Approved: true, User: common.User{Name: "rev4"}}
	rev5 := &common.Reviewer{Order: 5, Approved: false, User: common.User{Name: "rev5"}}
	rev6 := rev1
	rev6.Approved = false

	pr2 := common.PullRequest{Id: 116, Reviewers: []*common.Reviewer{rev4, rev5, &rev6}}
	repo.Stages = []Stage{{Reviewers: []*common.Reviewer{rev2, rev4, rev5, &rev6}}, {Reviewers: []*common.Reviewer{&rev1, rev3}}}

	repo.PRs = []common.PullRequest{pr1, pr2}

	busyReviewers := repo.getAssignedAndDoesNotApproveReviewers()

	assert.NotNil(t, busyReviewers)
	assert.True(t, busyReviewers.Cardinality() == 3)
	assert.True(t, busyReviewers.Contains(&rev6))
	assert.True(t, busyReviewers.Contains(rev3))
	assert.True(t, busyReviewers.Contains(rev5))
}

func Test_FindUserInReviewers_ExpectTheReviewer(t *testing.T) {

	repo := new(Repository)

	stage1 := Stage{Name: "TestStage1", Reviewers: []*common.Reviewer{getDummyReviewers()[0]}}
	stage2 := Stage{Name: "TestStage2", Reviewers: []*common.Reviewer{getDummyReviewers()[1], getDummyReviewers()[2]}}

	repo.Stages = []Stage{stage1, stage2}

	rv := repo.findReviewerByUsernameStage(getDummyReviewers()[2].User.Name)

	assert.NotNil(t, rv)
	assert.Equal(t, getDummyReviewers()[2], rv)
}

func Test_FindUserInReviewers_WithUnAvailableReviewerGivenUserInStages_ExpectNil(t *testing.T) {

	repo := new(Repository)

	stage1 := Stage{Name: "TestStage1", Reviewers: []*common.Reviewer{getDummyReviewers()[0]}}
	stage2 := Stage{Name: "TestStage2", Reviewers: []*common.Reviewer{getDummyReviewers()[1], getDummyReviewers()[2]}}

	repo.Stages = []Stage{stage1, stage2}

	rv := repo.findReviewerByUsernameStage(getDummyReviewers()[3].User.Name)

	assert.Nil(t, rv)
}

// AssignReviewersToPrs area
//When: 1 stage(Reviewer: 4 dummy reviewer, Policy: BYORDERINAVAILABLE), 1 PR(Owner: second-reviewer), BusyReviewers: 0
func TestRepository_AssignReviewersToPrs_ExpectFirstAsReviewerInStage(t *testing.T) {
	repo := new(Repository)

	stage := Stage{Name: "TestStage", Reviewers: getDummyReviewers(), Policy: BYORDERINAVAILABLE}

	pr := common.PullRequest{Id: 1903, Author: common.Author{User: getDummyReviewers()[1].User}}

	repo.Stages = []Stage{stage}
	repo.PRs = []common.PullRequest{pr}

	repo.AssignReviewersToPrs()

	assert.NotNil(t, repo.PRs)
	assert.True(t, len(repo.PRs) == 1)
	assert.NotNil(t, repo.PRs[0].Reviewers)
	assert.True(t, len(repo.PRs[0].Reviewers) == 1)
	assert.Equal(t, getDummyReviewers()[0], repo.PRs[0].Reviewers[0])
}

//When: 1 stage(Reviewer: 4 dummy reviewer, Policy: BYORDERINAVAILABLE), 1 PR(Owner: first-reviewer), BusyReviewers: 1(second reviewer)
func TestRepository_AssignReviewersToPrs_ExpectThirdInStage(t *testing.T) {
	repo := new(Repository)

	dummies := getDummyReviewers()

	owner := dummies[0].User
	reviewers := []*common.Reviewer{dummies[1]}

	stage := Stage{Name: "TestStage", Reviewers: dummies, Policy: BYORDERINAVAILABLE}

	pr1 := common.PullRequest{Id: 1903, Reviewers: reviewers, Author: common.Author{User: owner}}
	pr2 := common.PullRequest{Id: 116, Author: common.Author{User: owner}}

	repo.Stages = []Stage{stage}
	repo.PRs = []common.PullRequest{pr1, pr2}

	repo.AssignReviewersToPrs()

	assert.NotNil(t, repo.PRs)
	assert.True(t, len(repo.PRs) == 1)
	assert.NotNil(t, repo.PRs[0].Reviewers)
	assert.True(t, len(repo.PRs[0].Reviewers) == 1)
	assert.Equal(t, getDummyReviewers()[2], repo.PRs[0].Reviewers[0])
}

func TestRepository_AssignReviewersToPrs_With3Stage_2AvailableReviewerInFirstStageAnd1ReviewerInSecondAsAlsoOwner_ExpectFirstStagesReviewersToPR(t *testing.T) {
	repo := new(Repository)

	dummies := getDummyReviewers()

	owner := dummies[2]
	stage1Reviewers := []*common.Reviewer{dummies[0], dummies[1]}
	stage2Reviewers := []*common.Reviewer{owner}
	stage3Reviewers := []*common.Reviewer{dummies[0]}

	stage1 := Stage{Name: "TestStage1", Reviewers: stage1Reviewers, Policy: BYORDERINAVAILABLE}
	stage2 := Stage{Name: "TestStage2", Reviewers: stage2Reviewers, Policy: BYORDERINAVAILABLE}
	stage3 := Stage{Name: "TestStage3", Reviewers: stage3Reviewers, Policy: BYORDERINAVAILABLE}

	pr := common.PullRequest{Id: 116, Reviewers: nil, Author: common.Author{User: owner.User}}

	repo.Stages = []Stage{stage1, stage2, stage3}
	repo.PRs = []common.PullRequest{pr}

	repo.AssignReviewersToPrs()

	assert.NotNil(t, repo.PRs)
	assert.True(t, len(repo.PRs) == 1)
	assert.NotNil(t, repo.PRs[0].Reviewers)
	assert.True(t, len(repo.PRs[0].Reviewers) == 2)
	assert.Equal(t, getDummyReviewers()[0], repo.PRs[0].Reviewers[0])
	assert.Equal(t, getDummyReviewers()[1], repo.PRs[0].Reviewers[1])
}

func TestRepository_AssignReviewersToPrs_With3Stage_2AvailableReviewerInFirstStageAnd1ReviewerInSecondAsAlsoOwnerAndAvailableInTheLast_ExpectFirstStagesReviewersToPR(t *testing.T) {
	repo := new(Repository)

	dummies := getDummyReviewers()

	owner := dummies[2]
	stage1Reviewers := []*common.Reviewer{dummies[0], dummies[1]}
	stage2Reviewers := []*common.Reviewer{owner}
	stage3Reviewers := []*common.Reviewer{dummies[0], dummies[3]}

	stage1 := Stage{Name: "TestStage1", Reviewers: stage1Reviewers, Policy: BYORDERINAVAILABLE}
	stage2 := Stage{Name: "TestStage2", Reviewers: stage2Reviewers, Policy: BYORDERINAVAILABLE}
	stage3 := Stage{Name: "TestStage3", Reviewers: stage3Reviewers, Policy: BYORDERINAVAILABLE}

	pr := common.PullRequest{Id: 116, Reviewers: nil, Author: common.Author{User: owner.User}}

	repo.Stages = []Stage{stage1, stage2, stage3}
	repo.PRs = []common.PullRequest{pr}

	repo.AssignReviewersToPrs()

	assert.NotNil(t, repo.PRs)
	assert.True(t, len(repo.PRs) == 1)
	assert.NotNil(t, repo.PRs[0].Reviewers)
	assert.True(t, len(repo.PRs[0].Reviewers) == 3)
	assert.Equal(t, getDummyReviewers()[0], repo.PRs[0].Reviewers[0])
	assert.Equal(t, getDummyReviewers()[1], repo.PRs[0].Reviewers[1])
	assert.Equal(t, getDummyReviewers()[3], repo.PRs[0].Reviewers[2])
}

func TestRepository_AssignReviewersToPrs_With1StageHasOneReviewerAsOwnerAnd0AvailableReviewer_ExpectNoReviewersToPR(t *testing.T) {
	repo := new(Repository)

	dummies := getDummyReviewers()

	owner := dummies[2]
	stage1Reviewers := []*common.Reviewer{owner}

	stage1 := Stage{Name: "TestStage1", Reviewers: stage1Reviewers, Policy: BYORDERINAVAILABLE}

	pr := common.PullRequest{Id: 116, Reviewers: nil, Author: common.Author{User: owner.User}}

	repo.Stages = []Stage{stage1}
	repo.PRs = []common.PullRequest{pr}

	repo.AssignReviewersToPrs()

	assert.NotNil(t, repo.PRs)
	assert.True(t, len(repo.PRs) == 1)
	assert.Nil(t, repo.PRs[0].Reviewers)
}

func TestRepository_AssignReviewersToPrs_ExpectWithBusyReviewerInSecondStage(t *testing.T) {
	repo := new(Repository)

	dummies := getDummyReviewers()

	owner := dummies[0]

	stage1Reviewers := []*common.Reviewer{dummies[1]}
	stage2Reviewers := []*common.Reviewer{owner, dummies[2]}
	stage3Reviewers := []*common.Reviewer{dummies[3]}

	stage1 := Stage{Name: "TestStage1", Reviewers: stage1Reviewers, Policy: BYORDERINAVAILABLE}
	stage2 := Stage{Name: "TestStage2", Reviewers: stage2Reviewers, Policy: BYORDERINAVAILABLE}
	stage3 := Stage{Name: "TestStage3", Reviewers: stage3Reviewers, Policy: BYORDERINAVAILABLE}

	pr1Reviewers := []*common.Reviewer{dummies[2]}

	pr1 := common.PullRequest{Id: 1903, Reviewers: pr1Reviewers, Author: common.Author{User: dummies[3].User}}
	pr2 := common.PullRequest{Id: 116, Reviewers: nil, Author: common.Author{User: owner.User}}

	repo.Stages = []Stage{stage1, stage2, stage3}
	repo.PRs = []common.PullRequest{pr1, pr2}

	repo.AssignReviewersToPrs()

	assert.NotNil(t, repo.PRs)
	assert.True(t, len(repo.PRs) == 1)
	assert.NotNil(t, repo.PRs[0].Reviewers)
	assert.True(t, len(repo.PRs[0].Reviewers) == 3)
	assert.Equal(t, getDummyReviewers()[1], repo.PRs[0].Reviewers[0])
	assert.Equal(t, getDummyReviewers()[2], repo.PRs[0].Reviewers[1])
	assert.Equal(t, getDummyReviewers()[3], repo.PRs[0].Reviewers[2])

}

func TestRepository_AssignReviewersToPrs_ExpectAllReviewersOfSecondAndThirdStages(t *testing.T) {
	repo := new(Repository)

	dummies := getDummyReviewers()

	owner := dummies[0]

	stage1Reviewers := []*common.Reviewer{owner}
	stage2Reviewers := []*common.Reviewer{dummies[1], dummies[2]}
	stage3Reviewers := []*common.Reviewer{dummies[3], dummies[4]}

	stage1 := Stage{Name: "TestStage1", Reviewers: stage1Reviewers, Policy: BYORDERINAVAILABLE}
	stage2 := Stage{Name: "TestStage2", Reviewers: stage2Reviewers, Policy: BYORDERINAVAILABLE}
	stage3 := Stage{Name: "TestStage3", Reviewers: stage3Reviewers, Policy: BYORDERINAVAILABLE}

	pr := common.PullRequest{Id: 1903, Reviewers: nil, Author: common.Author{User: owner.User}}

	repo.Stages = []Stage{stage1, stage2, stage3}
	repo.PRs = []common.PullRequest{pr}

	repo.AssignReviewersToPrs()

	assert.NotNil(t, repo.PRs)
	assert.True(t, len(repo.PRs) == 1)
	assert.NotNil(t, repo.PRs[0].Reviewers)
	assert.True(t, len(repo.PRs[0].Reviewers) == 3)
	assert.Equal(t, getDummyReviewers()[1], repo.PRs[0].Reviewers[0])
	assert.Equal(t, getDummyReviewers()[2], repo.PRs[0].Reviewers[1])
	assert.Equal(t, getDummyReviewers()[3], repo.PRs[0].Reviewers[2])

}

func setMockEnvironmentVariables() {
	_ = os.Setenv("BITBUCKET_BASE_URL", "http://localhost/rest/api/1.0")
	_ = os.Setenv("BITBUCKET_PROJECT_PATH", "/projects/")
	_ = os.Setenv("BITBUCKET_REPOSITORY_PATH", "/repos/")
	_ = os.Setenv("BITBUCKET_PR_PATH", "/pull-requests")
	_ = os.Setenv("BITBUCKET_USER_PATH", "/permissions/users")
}
