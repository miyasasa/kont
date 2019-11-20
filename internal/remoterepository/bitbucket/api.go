package bitbucket

import (
	"bytes"
	"encoding/json"
	"kont/init/env"
	"kont/internal/client"
	"kont/internal/common"
	"kont/internal/repository"
	"log"
	"net/http"
	"strconv"
)

func fetchUsers(url string, start int) []common.User {
	req, _ := http.NewRequest("GET", url+"?start="+strconv.Itoa(start), nil)
	req.Header.Add("Authorization", env.BitbucketToken)

	page := common.UserPagination{}
	client.GET(req, &page)

	if page.IsLastPage {
		return page.GetUsers()
	}

	return append(fetchUsers(url, page.NextPageStart), page.GetUsers()...)
}

// ignored pagination request cause of pageSize equals 25
func fetchPRs(repo *repository.Repository) {
	req, _ := http.NewRequest("GET", repo.FetchPrsUrl, nil)
	req.Header.Add("Authorization", repo.Token)

	page := common.PRPagination{}
	client.GET(req, &page)

	repo.PRs = page.Values
}

func updatePRs(repo *repository.Repository) {

	for _, pr := range repo.PRs {

		url := repo.FetchPrsUrl + "/" + strconv.FormatInt(int64(pr.Id), 10)
		body, err := json.Marshal(pr)

		if err != nil {
			log.Printf("Pull-Request can not convert to json string")
		}

		req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(body))
		req.Header.Add("Authorization", repo.Token)
		req.Header.Add("Content-Type", "application/json")

		//client.PUT(req)

		log.Printf("%+v\n", pr)
	}
}
