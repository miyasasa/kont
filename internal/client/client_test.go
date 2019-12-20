package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"io/ioutil"
	"kont/internal/common"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestNewHttpClient_ExpectNewHttpClient(t *testing.T) {
	assertion := assert.New(t)

	d := NewHttpDispatcher()
	c := NewHttpClient(d)

	assertion.NotNil(c)
	assertion.NotNil(c.dispatcher)
}

func TestHttpClient_GET_Given200StatusCodeWithExpectedBody_ExpectToBindGivenInterface(t *testing.T) {
	assertion := assert.New(t)

	prUrl := "http://localhost:7990/rest/api/1.0/projects/BESG/repos/core-network/pull-requests?start=0"
	req, _ := http.NewRequest("GET", prUrl, nil)
	expectedPage := &common.PRPagination{Pagination: common.Pagination{Size: 10, Limit: 10, IsLastPage: true}}

	dispatcherMock := HttpDispatcherMock{}
	dispatcherMock.On("dispatch", req).Return(&http.Response{StatusCode: 200, Body: getBody(expectedPage)})

	c := NewHttpClient(dispatcherMock)

	page := common.PRPagination{}
	c.GET(req, &page)

	assertion.NotNil(page)
	assertion.Equal(10, page.Size)
	assertion.Equal(10, page.Limit)
	assertion.True(page.IsLastPage)
}

func TestHttpClient_GET_Given200StatusCodeWithUnExpectedBody_ExpectEmptyGivenInterfaceAndErrorLog(t *testing.T) {

	assertion := assert.New(t)
	prUrl := "http://localhost:7990/rest/api/1.0/projects/BESG/repos/core-network/pull-requests?start=0"
	req, _ := http.NewRequest("GET", prUrl, nil)
	expected := "{\"repo:\":\"core-network\"}"

	dispatcherMock := HttpDispatcherMock{}
	dispatcherMock.On("dispatch", req).Return(&http.Response{StatusCode: 200, Body: getBody(expected)})

	c := NewHttpClient(dispatcherMock)

	page := common.PRPagination{}
	output := captureOutput(func() {
		c.GET(req, &page)
	})

	assertion.NotNil(page)
	assertion.Equal(0, page.Size)
	assertion.Equal(0, page.Limit)
	assertion.False(page.IsLastPage)

	assertion.NotNil(output)
	assertion.Contains(output, fmt.Sprintf("client::GET::Response Body can not be converted to desired type-struct, Req-Url: %s", prUrl))
	assertion.Contains(output, "Error: json: cannot unmarshal string into Go value of type common.PRPagination")
}

func TestHttpClient_GET_Given404StatusCode_ExpectEmptyInterfaceAndErrorLog(t *testing.T) {
	assertion := assert.New(t)

	prUrl := "http://localhost:7990/rest/api/1.0/projects/BESG/repos/core-network1/pull-requests?start=0"
	req, _ := http.NewRequest("GET", prUrl, nil)
	expected := "{\"errors\": [{\"context\": null,\"message\": \"Repository BESG/core-network1 does not exist.\",\"exceptionName\": \"com.atlassian.bitbucket.repository.NoSuchRepositoryException\"}]}"

	dispatcherMock := HttpDispatcherMock{}
	dispatcherMock.On("dispatch", req).Return(&http.Response{StatusCode: 404, Body: getBody(expected)})

	c := NewHttpClient(dispatcherMock)

	page := common.PRPagination{}
	output := captureOutput(func() {
		c.GET(req, &page)
	})

	assertion.NotNil(page)
	assertion.Equal(0, page.Size)
	assertion.Equal(0, page.Limit)
	assertion.False(page.IsLastPage)

	assertion.NotNil(output)
	assertion.Contains(output, fmt.Sprintf("client::GET::Unexpected response; StatusCode: 404, Req-Url: %s", prUrl))
}

func TestHttpClient_GET_GivenNilByDispatcher_ExpectEmptyPrPagination(t *testing.T) {
	assertion := assert.New(t)

	prUrl := "http://localhost:7990/rest/api/1.0/projects/BESG/repos/core-network1/pull-requests?start=0"
	req, _ := http.NewRequest("GET", prUrl, nil)

	dispatcherMock := HttpDispatcherMock{}
	dispatcherMock.On("dispatch", req).Return(nil)

	c := NewHttpClient(dispatcherMock)

	page := common.PRPagination{}
	output := captureOutput(func() {
		c.GET(req, &page)
	})

	assertion.NotNil(page)
	assertion.Equal(0, page.Size)
	assertion.Equal(0, page.Limit)
	assertion.False(page.IsLastPage)

	assertion.Empty(output)
}

func TestHttpClient_PUT_Given200ResponseMock_NotExpectError(t *testing.T) {
	assertion := assert.New(t)

	prUrl := "http://localhost:7990/rest/api/1.0/projects/BESG/repos/core-network1/pull-requests/1903"
	req, _ := http.NewRequest("PUT", prUrl, nil)

	dispatcherMock := HttpDispatcherMock{}
	dispatcherMock.On("dispatch", req).Return(&http.Response{StatusCode: 200})

	c := NewHttpClient(dispatcherMock)
	output := captureOutput(func() {
		c.PUT(req)
	})

	assertion.Empty(output)

}

func TestHttpClient_PUT_Given500ResponseMock_ExpectErrorLog(t *testing.T) {
	assertion := assert.New(t)

	prUrl := "http://localhost:7990/rest/api/1.0/projects/BESG/repos/core-network1/pull-requests/1903"
	req, _ := http.NewRequest("PUT", prUrl, nil)

	dispatcherMock := HttpDispatcherMock{}
	dispatcherMock.On("dispatch", req).Return(&http.Response{StatusCode: 500})

	c := NewHttpClient(dispatcherMock)
	output := captureOutput(func() {
		c.PUT(req)
	})

	assertion.NotNil(output)
	assertion.Contains(output, fmt.Sprintf("client::PUT:: PR can not updated URL: %s StatusCode: 500", prUrl))
}

func TestHttpClient_PUT_GivenNilResponseMock_NotExpectError(t *testing.T) {
	assertion := assert.New(t)

	prUrl := "http://localhost:7990/rest/api/1.0/projects/BESG/repos/core-network1/pull-requests/1903"
	req, _ := http.NewRequest("PUT", prUrl, nil)

	dispatcherMock := HttpDispatcherMock{}
	dispatcherMock.On("dispatch", req).Return(nil)

	c := NewHttpClient(dispatcherMock)
	output := captureOutput(func() {
		c.PUT(req)
	})

	assertion.Empty(output)
}

type HttpDispatcherMock struct {
	mock.Mock
}

func (d HttpDispatcherMock) dispatch(req *http.Request) *http.Response {
	args := d.Called(req)

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(*http.Response)
}

func getBody(p interface{}) io.ReadCloser {
	requestByte, _ := json.Marshal(p)
	requestReader := bytes.NewReader(requestByte)

	return ioutil.NopCloser(requestReader)
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)

	return buf.String()
}
