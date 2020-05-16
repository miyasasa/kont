package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_httpDispatcher_Instance_ExpectNotEmpty(t *testing.T) {

	assert.NotNil(t, httpDispatcher)
}

func Test_HttpClientInstance_Instance_ExpectNotEmpty(t *testing.T) {

	assert.NotNil(t, HttpClientInstance)
	assert.NotNil(t, HttpClientInstance.dispatcher)
}
