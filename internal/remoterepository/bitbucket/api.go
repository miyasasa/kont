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

type Api interface {
	fetchUsers(url string, token string, start int) []common.User
	fetchPRs(repo *repository.Repository, start int)
	updatePRs(repo *repository.Repository)
}

type BitbucketApi struct {
	client client.Client
}

func NewBitbucketApi(client client.Client) *BitbucketApi {
	return &BitbucketApi{client: client}
}

func (b *BitbucketApi) fetchUsers(url string, token string, start int) []common.User {
	req, _ := http.NewRequest("GET", url+"?start="+strconv.Itoa(start), nil)
	req.Header.Add("Authorization", token)

	page := common.UserPagination{}
	b.client.HandleToInterface(req, &page)

	if page.IsLastPage || page.NextPageStart == 0 {
		return page.GetUsers()
	}

	return append(b.fetchUsers(url, token, page.NextPageStart), page.GetUsers()...)
}

func (b *BitbucketApi) fetchPRs(repo *repository.Repository, start int) {
	req, _ := http.NewRequest("GET", repo.FetchPrsUrl+"?start="+strconv.Itoa(start), nil)
	req.Header.Add("Authorization", repo.Token)

	page := common.PRPagination{}
	b.client.HandleToInterface(req, &page)

	repo.PRs = append(repo.PRs, page.Values...)

	if !page.IsLastPage && page.NextPageStart != 0 {
		b.fetchPRs(repo, page.NextPageStart)
	}

}

func (b *BitbucketApi) updatePRs(repo *repository.Repository) {

	for i := range repo.PRs {
		go b.addReviewersToPR(i, repo)

		if len(repo.DefaultComment) > 0 {
			go b.addDefaultCommentToPR(i, repo)
		}
	}
}

func (b *BitbucketApi) addReviewersToPR(i int, repo *repository.Repository) {
	pr := repo.PRs[i]
	url := repo.FetchPrsUrl + "/" + strconv.FormatInt(int64(pr.Id), 10)

	uPR := MapPullRequestToUpdateModel(pr)
	body, err := json.Marshal(uPR)

	if err != nil {
		log.Printf("api::addReviewersToPR, Pull-Request can not convert to json string")
	}

	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	req.Header.Add("Authorization", repo.Token)
	req.Header.Add("Content-Type", "application/json")

	b.client.Handle(req)

	log.Printf("%+v\n", uPR)
}

func (b *BitbucketApi) addDefaultCommentToPR(i int, repo *repository.Repository) {
	pr := repo.PRs[i]
	url := repo.FetchPrsUrl + "/" + strconv.FormatInt(int64(pr.Id), 10) + "/comments"

	dC := NewPullRequestDefaultCommentUpdateModel(repo.DefaultComment, pr.Author.User.Name)
	body, err := json.Marshal(dC)

	if err != nil {
		log.Printf("api::addDefaultCommentToPR, default-comment can not convert to json string")
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Add("Authorization", repo.Token)
	req.Header.Add("Content-Type", "application/json")

	b.client.Handle(req)

	log.Printf("Added default-comment to %d, Default-Comment: %s \n", pr.Id, dC)
}
