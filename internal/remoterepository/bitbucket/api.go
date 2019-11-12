package bitbucket

import (
	"bytes"
	"encoding/json"
	"log"
	"miya/internal/client"
	"miya/internal/common"
	"miya/internal/repository"
	"net/http"
	"strconv"
)

// ignored pagination request cause of pageSize equals 25
func fetchPRs(repo *repository.Repository) {
	req, _ := http.NewRequest("GET", repo.FetchPrURL, nil)
	req.Header.Add("Authorization", repo.Token)

	page := common.Pagination{}
	client.GET(req, &page)

	repo.PR = page.Values
}

func updatePRs(repo *repository.Repository) {

	for _, pr := range repo.PR {

		url := repo.FetchPrURL + "/" + strconv.FormatInt(int64(pr.Id), 10)
		body, err := json.Marshal(pr)

		if err != nil {
			log.Printf("Pull-Request can not convert to json string")
		}

		req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(body))
		req.Header.Add("Authorization", repo.Token)
		req.Header.Add("Content-Type", "application/json")

		client.PUT(req)

		log.Printf("%+v\n", pr)
	}
}