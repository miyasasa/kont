package bitbucket

import (
	"kont/internal/common"
	"kont/internal/repository"
)

func AppendPullRequestToRepo(pullRequests []PullRequest, repo *repository.Repository) {
	for _, pl := range pullRequests {
		pr := mapPullRequestToCommonPR(pl, repo)
		repo.PRs = append(repo.PRs, pr)
	}
}

func mapPullRequestToCommonPR(pullRequest PullRequest, repo *repository.Repository) common.PR {
	pr := common.PR{}

	pr.Id = pullRequest.Id
	pr.Version = pullRequest.Version
	pr.Title = pullRequest.Title
	pr.Description = pullRequest.Description
	pr.Reviewers = mapReviewersToCommonReviewers(pullRequest.Reviewers, repo)
	pr.Author = mapAuthorToCommonAuthor(pullRequest.Author)

	return pr
}

func mapAuthorToCommonAuthor(author Author) common.Author {
	au := common.Author{}
	au.User = mapUserToCommonUSer(author.User)

	return au
}

func mapUserToCommonUSer(user User) common.User {
	u := common.User{}
	u.Name = user.Name
	u.DisplayName = user.DisplayName

	return u
}

func mapReviewersToCommonReviewers(reviewers []*Reviewer, repo *repository.Repository) []*common.Reviewer {
	rvs := make([]*common.Reviewer, 0)

	for _, rv := range reviewers {
		r := mapReviewerToCommonReviewer(rv, repo)
		rvs = append(rvs, r)
	}

	return rvs
}

func mapReviewerToCommonReviewer(reviewer *Reviewer, repo *repository.Repository) *common.Reviewer {
	for _, s := range repo.Stages {
		if reviewer := s.GetReviewerByUserName(reviewer.User.Name); reviewer != nil {
			return reviewer
		}
	}

	return nil
}
