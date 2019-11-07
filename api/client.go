package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func Send(req *http.Request, i interface{}) interface{} {

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("An Error occur sen http-request to ... %s", err)
	} else {
		err := json.NewDecoder(resp.Body).Decode(i)
		if err != nil {
			log.Fatalf("Response Body can not converted desired type-struct ...req: %s, %s", req.URL, err)
		}
	}

	return i
}
