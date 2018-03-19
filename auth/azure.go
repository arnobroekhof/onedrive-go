package auth

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	keysUrl = "https://login.microsoftonline.com/common/discovery/v2.0/keys"
)

// Azure struct
//Grab the certificate (the first value in the x5c property array) from either https://login.microsoftonline.com/common/discovery/keys or https://login.microsoftonline.com/common/discovery/v2.0/keys, matching kid and x5t from the id_token.
//Wrap the certificate in -----BEGIN CERTIFICATE-----\n and \n-----END CERTIFICATE----- (the newlines seem to matter), and use the result as Public Key (in conjunction with the id_token, on https://jwt.io/ ).
type Azure struct {
	Keys []struct {
		Kty    string   `json:"kty"`
		Use    string   `json:"use"`
		Kid    string   `json:"kid"`
		X5T    string   `json:"x5t"`
		N      string   `json:"n"`
		E      string   `json:"e"`
		X5C    []string `json:"x5c"`
		Issuer string   `json:"issuer"`
	} `json:"keys"`
}

func RetrieveKeys() (*Azure, error) {

	// create httpClient with timeout option
	var httpClient = &http.Client{Timeout: 10 * time.Second}
	res, err := httpClient.Get(keysUrl)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	var azure *Azure
	mErr := json.Unmarshal(body, &azure)
	if err != nil {
		return nil, mErr
	}

	return azure, nil
}

func (a Azure) GetX5TMatchingPubKey(x5t string) (string, error) {
	for _, key := range a.Keys {
		if key.X5T == x5t {
			return addCertificateHeaderAndFooter(key.X5C[0]), nil
		}
	}
	return "", errors.New("no matching x5t key found")
}

func (a Azure) GetKIDMatchingPubKey(kid string) (string, error) {
	for _, key := range a.Keys {
		if key.Kid == kid {
			return addCertificateHeaderAndFooter(key.X5C[0]), nil
		}
	}
	return "", errors.New("no matching x5t key found")
}

func addCertificateHeaderAndFooter(key string) string {
	headerString := "-----BEGIN CERTIFICATE-----"
	footerString := "-----END CERTIFICATE-----"
	pubKey := []string{headerString, key, footerString}
	return strings.Join(pubKey, "\n")
}

func ValidateToken(tokenString string) (authenticated bool, claimMap map[string]interface{}, err error) {

	if tokenString == "" {
		return authenticated, claimMap, errors.New("token cannot be empty")
	}

	var token *jwt.Token
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		azure, err := RetrieveKeys()
		if err != nil {
			return nil, err
		}

		x5t := token.Header["x5t"].(string)
		stringKey, err := azure.GetX5TMatchingPubKey(x5t)

		var verifyKey *rsa.PublicKey
		verifyKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(stringKey))
		if err != nil {
			return nil, err
		}

		return verifyKey, nil
	})

	if err != nil {
		return false, nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token.Valid, token.Claims.(jwt.MapClaims), nil
	} else {
		return token.Valid, nil, err
	}
}
