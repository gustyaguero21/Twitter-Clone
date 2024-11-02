package utils

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCreateResponse(t *testing.T) {
	//given
	status := 200
	message := "Success"

	//act
	response := CreateResponse(status, message)

	//asserts
	assert.Equal(t, status, response.Status)
}
