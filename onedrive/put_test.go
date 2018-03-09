package onedrive

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOnedrive_PutWithoutPath(t *testing.T) {
	od, err := NewOnedrive("fake-token")
	assert.IsType(t, &Onedrive{}, od)
	assert.Nil(t, err)
	resp, pErr := od.Put("")
	assert.Error(t, pErr, "path cannot be empty")
	assert.Empty(t, resp)
}

func TestOnedrive_PutWithPath(t *testing.T) {
	od, err := NewOnedrive("fake-token")
	assert.IsType(t, &Onedrive{}, od)
	assert.Nil(t, err)
	resp, pErr := od.Put("/test/file")
	assert.Nil(t, pErr)
	assert.NotEmpty(t, resp)
}
