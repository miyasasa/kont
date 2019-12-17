package client

import (
	"bytes"
	"fmt"
	"github.com/jarcoal/httpmock"
	assert2 "github.com/stretchr/testify/assert"
	"kont/internal/common"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestGET_Given200StatusCodeWithExpectedBody_ExpectToBindGivenInterface(t *testing.T) {

	prUrl := "http://localhost:7990/rest/api/1.0/projects/BESG/repos/core-network/pull-requests?start=0"
	assert := assert2.New(t)

	//active http mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	//simulate response
	expectedPage := common.PRPagination{Pagination: common.Pagination{Size: 10, Limit: 10, IsLastPage: true}}
	responder, _ := httpmock.NewJsonResponder(200, &expectedPage)

	httpmock.RegisterResponder("GET", prUrl, responder)

	//get expected response
	page := common.PRPagination{}
	req, _ := http.NewRequest("GET", prUrl, nil)
	GET(req, &page)

	assert.NotNil(page)
	assert.Equal(expectedPage.Size, page.Size)
	assert.Equal(expectedPage.Limit, page.Limit)
	assert.True(page.IsLastPage)
}

func TestGET_Given200StatusCodeWithUnExpectedBody_ExpectEmptyGivenInterfaceAndErrorLog(t *testing.T) {

	prUrl := "http://localhost:7990/rest/api/1.0/projects/BESG/repos/core-network/pull-requests?start=0"
	assert := assert2.New(t)

	//active http mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	//simulate response
	responder, _ := httpmock.NewJsonResponder(200, "{\"repo:\":\"core-network\"}")

	httpmock.RegisterResponder("GET", prUrl, responder)

	//get expected response
	page := common.PRPagination{}
	req, _ := http.NewRequest("GET", prUrl, nil)
	output := captureOutput(func() {
		GET(req, &page)
	})

	assert.NotNil(page)
	assert.Equal(0, page.Size)
	assert.Equal(0, page.Limit)
	assert.False(page.IsLastPage)

	assert.NotNil(output)
	assert.Contains(output, fmt.Sprintf("client::GET::Response Body can not be converted to desired type-struct, Req-Url: %s", prUrl))
	assert.Contains(output, "Error: json: cannot unmarshal string into Go value of type common.PRPagination")
}

func TestGET_Given404_ExpectEmptyInterfaceAndErrorLog(t *testing.T) {

	prUrl := "http://localhost:7990/rest/api/1.0/projects/BESG/repos/core-network1/pull-requests?start=0"
	assert := assert2.New(t)

	//active http mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	//simulate response
	responder, _ := httpmock.NewJsonResponder(404, "{\"errors\": [{\"context\": null,\"message\": \"Repository BESG/core-network1 does not exist.\",\"exceptionName\": \"com.atlassian.bitbucket.repository.NoSuchRepositoryException\"}]}")

	httpmock.RegisterResponder("GET", prUrl, responder)

	//get expected response
	page := common.PRPagination{}
	req, _ := http.NewRequest("GET", prUrl, nil)

	output := captureOutput(func() {
		GET(req, &page)
	})

	assert.NotNil(page)
	assert.Equal(0, page.Size)
	assert.Equal(0, page.Limit)
	assert.False(page.IsLastPage)

	assert.NotNil(output)
	assert.Contains(output, fmt.Sprintf("client::GET::Unexpected response; StatusCode: 404, Req-Url: %s", prUrl))
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)

	return buf.String()
}
