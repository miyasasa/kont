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
	c := client.NewHttpClient(client.NewHttpDispatcher())
	c.HandleToInterface(req, &page)

	if page.IsLastPage || page.NextPageStart == 0 {
		return page.GetUsers()
	}

	return append(fetchUsers(url, token, page.NextPageStart), page.GetUsers()...)
}

func fetchPRs(repo *repository.Repository, start int) {
	req, _ := http.NewRequest("GET", repo.FetchPrsUrl+"?start="+strconv.Itoa(start), nil)
	req.Header.Add("Authorization", repo.Token)

	page := common.PRPagination{}
	c := client.NewHttpClient(client.NewHttpDispatcher())
	c.HandleToInterface(req, &page)

	repo.PRs = append(repo.PRs, page.Values...)

	if !page.IsLastPage && page.NextPageStart != 0 {
		fetchPRs(repo, page.NextPageStart)
	}

}

func updatePRs(repo *repository.Repository) {

	for _, pr := range repo.PRs {
		go addReviewersToPR(pr, repo)
		go addDefaultCommentToPR(pr, repo)
	}
}

func addReviewersToPR(pr common.PullRequest, repo *repository.Repository) {
	url := repo.FetchPrsUrl + "/" + strconv.FormatInt(int64(pr.Id), 10)

	uPR := MapPullRequestToUpdateModel(pr)
	body, err := json.Marshal(uPR)

	if err != nil {
		log.Printf("api::addReviewersToPR, Pull-Request can not convert to json string")
	}

	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	req.Header.Add("Authorization", repo.Token)
	req.Header.Add("Content-Type", "application/json")

	c := client.NewHttpClient(client.NewHttpDispatcher())
	c.Handle(req)

	log.Printf("%+v\n", pr)
}

func addDefaultCommentToPR(pr common.PullRequest, repo *repository.Repository) {
	url := repo.FetchPrsUrl + "/" + strconv.FormatInt(int64(pr.Id), 10) + "/comments"

	dC := NewPullRequestDefaultCommentUpdateModel(repo.DefaultComment, pr.Author.User.Name)
	body, err := json.Marshal(dC)

	if err != nil {
		log.Printf("api::addDefaultCommentToPR, default-comment can not convert to json string")
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Add("Authorization", repo.Token)
	req.Header.Add("Content-Type", "application/json")

	c := client.NewHttpClient(client.NewHttpDispatcher())
	c.Handle(req)

	log.Printf("Added default-comment to %d, Default-Comment: %s \n", pr.Id, dC)
}
