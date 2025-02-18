package errorx

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromError(t *testing.T) {
	excepted := NewStatCodeError(400, 1, "not found")
	err := WithStack(excepted)

	c, ok := FromError(err)
	assert.True(t, ok)
	assert.Equal(t, excepted, c)
}

func TestFromErrorFail(t *testing.T) {
	err := errors.New("Hello ")

	c, ok := FromError(err)
	assert.False(t, ok)
	assert.Nil(t, c)
}
