package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestNewHttpDispatcher_ExpectNewDispatcher(t *testing.T) {
	assertion := assert.New(t)

	d := NewHttpDispatcher()

	assertion.NotNil(d)
}

func TestGET_Given200ResponseToResponder_Expect200AsResponse(t *testing.T) {
	assertion := assert.New(t)

	//active http mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	//simulate response
	prUrl := "http://localhost:7990/rest/api/1.0/projects/BESG/repos/core-network1/pull-requests?start=0"
	body := "Güven Yalçın"
	responder := httpmock.NewStringResponder(200, body)

	httpmock.RegisterResponder("PUT", prUrl, responder)

	d := HttpDispatcher{}

	//get expected response
	req, _ := http.NewRequest("PUT", prUrl, nil)
	resp := d.dispatch(req)

	assertion.NotNil(resp)
	assertion.Equal(200, resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assertion.Nil(err)
	assertion.Equal(body, string(bodyBytes))

}

func TestGET_Given404ResponseToResponder_Expect404AsResponse(t *testing.T) {
	assertion := assert.New(t)

	//active http mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	//simulate response
	prUrl := "http://localhost:7990/rest/api/1.0/projects/BESG/repos/core-network1/pull-requests?start=0"
	body := `{"errors": [{"context": null,"message": "Repository BESG/core-network1 does not exist.","exceptionName": "com.atlassian.bitbucket.repository.NoSuchRepositoryException"}]}`

	responder, _ := httpmock.NewJsonResponder(404, body)
	httpmock.RegisterResponder("PUT", prUrl, responder)

	d := HttpDispatcher{}

	//get expected response
	req, _ := http.NewRequest("PUT", prUrl, nil)
	resp := d.dispatch(req)

	assertion.NotNil(resp)
	assertion.Equal(404, resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assertion.Nil(err)

	var result string
	err = json.Unmarshal(bodyBytes, &result)
	assertion.Nil(err)

	assertion.Equal(body, result)
}

func TestGET_GivenNilResponseToResponder_ExpectNilAsResponseAndErrorLogOnConsole(t *testing.T) {
	assertion := assert.New(t)

	//active http mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	//simulate response
	prUrl := "http://localhost:7990/rest/api/1.0/projects/BESG/repos/core-network1/pull-requests?start=0"
	req, _ := http.NewRequest("PUT", prUrl, nil)

	responder := func(req *http.Request) (*http.Response, error) {
		return nil, errors.New("request error")
	}

	httpmock.RegisterResponder("PUT", prUrl, responder)

	d := HttpDispatcher{}

	resp := &http.Response{}
	assertion.NotNil(resp)

	output := captureOutput(func() {
		resp = d.dispatch(req)
	})

	assertion.Nil(resp)
	assertion.Contains(output, fmt.Sprintf("Dispatcher::dispatch An Error occur send http-request to:  Put %s: request error", prUrl))
}
