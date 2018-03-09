package onedrive

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOnedrive_GetWithoutPath(t *testing.T) {
	od, err := NewOnedrive("fake-token")
	assert.IsType(t, &Onedrive{}, od)
	assert.Nil(t, err)
	resp, pErr := od.Get("")
	assert.Error(t, pErr, "path cannot be empty")
	assert.Empty(t, resp)
}

func TestOnedrive_GetWithPath(t *testing.T) {
	od, err := NewOnedrive("fake-token")
	assert.IsType(t, &Onedrive{}, od)
	assert.Nil(t, err)
	resp, pErr := od.Get("/test/file")
	assert.Nil(t, pErr)
	assert.IsType(t, []byte{}, resp)
}
