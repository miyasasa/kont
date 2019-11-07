package bitbucket

import (
	"miya/api"
	"miya/api/bitbucket/model"
	"miya/init/env"
	"net/http"
)

// ignored pagination request cause of pageSize equals 25
func fetchPRs() []model.PullRequest {
	req, _ := http.NewRequest("GET", env.BitbucketFetchPrListUrl, nil)
	req.Header.Add("Authorization", env.BitbucketToken)

	page := model.Pagination{}
	api.Send(req, &page)

	return page.Values
}
