package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler_JsonError(t *testing.T) {
	msg := "Hello Json"
	result := JsonError(msg)
	assert.Equal(t, []byte(`{"message":"Hello Json"}`), result)
}
