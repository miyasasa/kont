package client

import (
	"encoding/json"
	"log"
	"net/http"
)

type Client interface {
	GET(req *http.Request, i interface{})
	PUT(req *http.Request)
}

type HttpClient struct {
	dispatcher Dispatcher
}

func NewHttpClient(dispatcher Dispatcher) *HttpClient {
	return &HttpClient{dispatcher: dispatcher}
}

func (c *HttpClient) GET(req *http.Request, i interface{}) {

	resp := c.dispatcher.dispatch(req)

	if resp != nil {
		if resp.StatusCode == 200 {
			err := json.NewDecoder(resp.Body).Decode(i)
			if err != nil {
				log.Printf("client::GET::Response Body can not be converted to desired type-struct, Req-Url: %s, Error: %s", req.URL, err)
			}
		} else {
			log.Printf("client::GET::Unexpected response; StatusCode: %d, Req-Url: %v", resp.StatusCode, req.URL)
		}
	}
}

func (c HttpClient) PUT(req *http.Request) {
	resp := c.dispatcher.dispatch(req)

	if resp != nil && resp.StatusCode != 200 {
		log.Printf("client::PUT:: PR can not updated URL: %s StatusCode: %v", req.URL, resp.StatusCode)
	}
}
