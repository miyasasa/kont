package bitbucket

import (
	"miya/init/env"
	"miya/internal/client"
	"miya/internal/common"
	"net/http"
)

// ignored pagination request cause of pageSize equals 25
func fetchPRs() []common.PullRequest {
	req, _ := http.NewRequest("GET", env.BitbucketFetchPrListUrl, nil)
	req.Header.Add("Authorization", env.BitbucketToken)

	page := common.Pagination{}
	client.Send(req, &page)

	return page.Values
}
