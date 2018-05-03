package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRetrieveOpenIdConfiguration(t *testing.T) {
	openIdconfig, err := retrieveOpenIdConfiguration()
	assert.NoError(t, err, "error expected nil")
	assert.IsType(t, &openIDConfiguration{}, openIdconfig)

	keys, err := openIdconfig.RetrieveJwksKeys()
	assert.NoError(t, err)
	assert.IsType(t, &jwksKeys{}, keys)
	assert.Len(t, keys.Keys, 5)
}

func TestInvalidIdToken(t *testing.T) {
	authenticated, claims, err := ValidateIdToken("12345678910")
	assert.Error(t, err, "token contains an invalid number of segments")
	assert.Nil(t, claims, "Claims cannot be nil")
	assert.False(t, authenticated)
}

func TestNoIdTokenGiven(t *testing.T) {
	authenticated, claims, err := ValidateIdToken("")
	assert.Error(t, err, "no id token provided")
	assert.Nil(t, claims, "Claims cannot be nil")
	assert.False(t, authenticated)
}

func TestGetClaimsFromAccessToken(t *testing.T) {
	claims, err := GetClaimsFromAccessToken("")
	assert.Error(t, err, "no access token provided")
	assert.Nil(t, claims)

}
