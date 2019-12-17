package client

import (
	"encoding/json"
	"log"
	"net/http"
)

func GET(req *http.Request, i interface{}) interface{} {

	resp := send(req)

	if resp != nil && resp.StatusCode == 200 {
		err := json.NewDecoder(resp.Body).Decode(i)
		if err != nil {
			log.Printf("client::GET::Response Body can not be converted to desired type-struct, Req-Url: %s, Error: %s", req.URL, err)
		}
	} else {
		log.Printf("client::GET::Unexpected response; StatusCode: %d, Req-Url: %s", resp.StatusCode, req.URL)
	}

	return i
}

func PUT(req *http.Request) {
	resp := send(req)

	if resp.StatusCode != 200 {
		log.Printf("client::PUT:: PR can not updated URL: %s StatusCode: %v", req.URL, resp.StatusCode)
	}
}

func send(req *http.Request) *http.Response {
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("An Error occur sen http-request to ... %s", err)
		return nil
	}

	return resp
}
