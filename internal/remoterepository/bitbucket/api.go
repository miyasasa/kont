package bitbucket

import (
	"bytes"
	"encoding/json"
	"kont/internal/client"
	"kont/internal/common"
	"kont/internal/repository"
	"log"
	"net/http"
	"strconv"
)

func fetchUsers(url string, token string, start int) []common.User {
	req, _ := http.NewRequest("GET", url+"?start="+strconv.Itoa(start), nil)
	req.Header.Add("Authorization", token)

	page := common.UserPagination{}
	client.GET(req, &page)

	if page.IsLastPage || page.NextPageStart == 0 {
		return page.GetUsers()
	}

	return append(fetchUsers(url, token, page.NextPageStart), page.GetUsers()...)
}

func fetchPRs(repo *repository.Repository, start int) {
	req, _ := http.NewRequest("GET", repo.FetchPrsUrl+"?start="+strconv.Itoa(start), nil)
	req.Header.Add("Authorization", repo.Token)

	page := common.PRPagination{}
	client.GET(req, &page)

	repo.PRs = append(repo.PRs, page.Values...)

	if !page.IsLastPage && page.NextPageStart != 0 {
		fetchPRs(repo, page.NextPageStart)
	}

}

func updatePRs(repo *repository.Repository) {

	for _, pr := range repo.PRs {

		url := repo.FetchPrsUrl + "/" + strconv.FormatInt(int64(pr.Id), 10)

		uPR := MapPullRequestToUpdateModel(pr)
		body, err := json.Marshal(uPR)

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
