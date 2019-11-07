package bitbucket

import (
	"fmt"
	"log"
	"miya/api"
	"miya/api/bitbucket/model"
	"miya/init/env"
	"net/http"
)

func Listen() {
	fmt.Println("Bitbucket-PR is listening....")
	fetchPRs()
}

func fetchPRs() {
	req, _ := http.NewRequest("GET", env.BitbucketFetchPrListUrl, nil)
	req.Header.Add("Authorization", env.BitbucketToken)

	page := model.Pagination{}
	api.Send(req, &page)
	log.Printf("Response... %s", page.ToString())
}
