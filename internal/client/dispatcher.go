package client

import (
	"log"
	"net/http"
)

type Dispatcher interface {
	dispatch(req *http.Request) *http.Response
}

type HttpDispatcher struct{}

func NewHttpDispatcher() *HttpDispatcher {
	return &HttpDispatcher{}
}

func (d HttpDispatcher) dispatch(req *http.Request) *http.Response {
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Dispatcher::dispatch An Error occur send http-request to:  %s", err)
		return nil
	}

	return resp
}
