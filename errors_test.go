package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.Error(t, New("new error"))
	assert.Contains(t, New("new error").Error(), "new error")
}

func TestNewf(t *testing.T) {
	assert.Error(t, Newf("new error"))
	assert.Contains(t, Newf("this is new error").Error(), "this is new error")
	assert.Contains(t, Newf("%s %s", "new", "error").Error(), "new error")
}

func TestWrap(t *testing.T) {
	err := Wrap(errors.New("standard pkg error"), "wrap new error")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "standard pkg error")
	assert.Contains(t, err.Error(), "wrap new error")
}

func TestWrapf(t *testing.T) {
	assert.Error(t, Wrapf(errors.New("standard pkg error"), "wrap new error"))
	err := Wrapf(errors.New("standard pkg error"), "wrap %s %s", "new", "error")
	assert.Contains(t, err.Error(), "standard pkg error")
	assert.Contains(t, err.Error(), "wrap new error")
}
