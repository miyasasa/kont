package client

var httpDispatcher = NewHttpDispatcher()

var HttpClientInstance = NewHttpClient(httpDispatcher)
