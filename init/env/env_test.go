package env

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetServerPortVariableWithDefaultValue(t *testing.T) {

	assert.NotEmpty(t, ServerPort)
	assert.Equal(t, "1903", ServerPort)
}

func TestGivenSERVER_PORTEnvAs8080_ExpectSERVER_PORTEquals8080_AndServerPortVariableIsStill1903(t *testing.T) {

	assert.NotEmpty(t, ServerPort)
	assert.Equal(t, "1903", ServerPort)

	err := os.Setenv("SERVER_PORT", "8080")
	assert.Nil(t, err)

	serverPort := getEnv("SERVER_PORT", "1234")

	assert.NotEmpty(t, serverPort)
	assert.Equal(t, "8080", serverPort)
	assert.Equal(t, "1903", ServerPort)
}
