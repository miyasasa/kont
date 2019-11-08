package client

import (
	"encoding/json"
	"log"
	"net/http"
)

func GET(req *http.Request, i interface{}) interface{} {

	resp := send(req)

	if resp != nil {
		err := json.NewDecoder(resp.Body).Decode(i)
		if err != nil {
			log.Fatalf("client::GET::Response Body can not converted desired type-struct ...req: %s, %s", req.URL, err)
		}
	}

	return i
}

func PUT(req *http.Request) {
	resp := send(req)

	if resp.StatusCode != 200 {
		log.Fatalf("client::PUT:: PR can not updated URL: %s StatusCode: %v", req.URL, resp.StatusCode)
	}
}

func send(req *http.Request) *http.Response {
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("An Error occur sen http-request to ... %s", err)
		return nil
	}

	return resp
}
