package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRetrieveKeys(t *testing.T) {
	azure := retrieveKeys(t)

	assert.NotEmpty(t, azure)
	t.Logf("retrieved %v keys", len(azure.Keys))
}
func retrieveKeys(t *testing.T) *Azure {
	azure, err := RetrieveKeys()
	assert.Nil(t, err)
	return azure
}

func TestAzure_GetX5TMatchingPubKey(t *testing.T) {
	azure := retrieveKeys(t)

	assert.NotEmpty(t, azure)

	pubKey, err := azure.GetX5TMatchingPubKey("SSQdhI1cKvhQEDSJxE2gGYs40Q0")
	certificateChecking(t, err, pubKey)
}
func certificateChecking(t *testing.T, err error, pubKey string) {
	assert.Nil(t, err)
	assert.NotEmpty(t, pubKey)
	assert.IsType(t, "", pubKey)
	assert.Contains(t, pubKey, "-----BEGIN CERTIFICATE-----")
	assert.Contains(t, pubKey, "-----END CERTIFICATE-----")
}

func TestAzure_GetKIDMatchingPubKey(t *testing.T) {
	azure := retrieveKeys(t)

	assert.NotEmpty(t, azure)
	t.Logf("retrieved %v keys", len(azure.Keys))

	pubKey, err := azure.GetX5TMatchingPubKey("SSQdhI1cKvhQEDSJxE2gGYs40Q0")
	certificateChecking(t, err, pubKey)
}

func TestValidateToken(t *testing.T) {
	valid, claims, err := ValidateToken("")
	assert.Equal(t, false, valid)
	assert.Equal(t, map[string]interface{}(nil), claims)
	assert.EqualError(t, err, "token cannot be empty")
}

func TestInvalidToken(t *testing.T) {
	valid, claims, err := ValidateToken("1234567890")
	assert.Equal(t, false, valid)
	assert.Equal(t, map[string]interface{}(nil), claims)
	assert.EqualError(t, err, "token contains an invalid number of segments")
}
