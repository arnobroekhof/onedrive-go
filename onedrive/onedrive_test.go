package onedrive

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOnedriveWithToken(t *testing.T) {

	od, err := NewOnedrive("fake-token")
	assert.IsType(t, &Onedrive{}, od)
	assert.Nil(t, err)

}

func TestNewOnedriveWithoutToken(t *testing.T) {
	od, err := NewOnedrive("")
	assert.Nil(t, od)
	assert.Error(t, err, "token cannot be empty")
}

func TestReadGraphURL(t *testing.T) {
	od, err := NewOnedrive("fake-token")
	assert.IsType(t, &Onedrive{}, od)
	assert.Nil(t, err)
	t.Logf("Using drive url: %s", constructOneDriveUrl())
	assert.Equal(t, constructOneDriveUrl(), od.graphUrl)
}

func TestOneDrive_Connect(t *testing.T) {

}
