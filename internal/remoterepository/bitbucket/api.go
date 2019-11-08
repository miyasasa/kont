package bitbucket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"miya/internal/client"
	"miya/internal/common"
	"miya/internal/repository"
	"net/http"
	"strconv"
)

// ignored pagination request cause of pageSize equals 25
// return repo ?
func fetchPRs(repo repository.Repository) []common.PullRequest {
	req, _ := http.NewRequest("GET", repo.FetchPrURL, nil)
	req.Header.Add("Authorization", repo.Token)

	page := common.Pagination{}
	client.GET(req, &page)

	return page.Values
}

func updatePRsForAddingReviewers(repo repository.Repository) {

	for _, pr := range repo.PR {

		url := repo.FetchPrURL + "/" + strconv.FormatInt(int64(pr.Id), 10)
		body, err := json.Marshal(pr)

		if err != nil {
			log.Printf("Pull-Request can not convert to json string")
		}

		req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(body))
		req.Header.Add("Authorization", repo.Token)
		req.Header.Add("Content-Type", "application/json")

		fmt.Printf("%+v\n", pr)

		client.PUT(req)
	}
}
