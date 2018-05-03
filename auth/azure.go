package auth

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
	"time"
)

type openIDConfiguration struct {
	AuthorizationEndpoint             string      `json:"authorization_endpoint"`
	TokenEndpoint                     string      `json:"token_endpoint"`
	TokenEndpointAuthMethodsSupported []string    `json:"token_endpoint_auth_methods_supported"`
	JwksURI                           string      `json:"jwks_uri"`
	ResponseModesSupported            []string    `json:"response_modes_supported"`
	SubjectTypesSupported             []string    `json:"subject_types_supported"`
	IDTokenSigningAlgValuesSupported  []string    `json:"id_token_signing_alg_values_supported"`
	HTTPLogoutSupported               bool        `json:"http_logout_supported"`
	FrontchannelLogoutSupported       bool        `json:"frontchannel_logout_supported"`
	EndSessionEndpoint                string      `json:"end_session_endpoint"`
	ResponseTypesSupported            []string    `json:"response_types_supported"`
	ScopesSupported                   []string    `json:"scopes_supported"`
	Issuer                            string      `json:"issuer"`
	ClaimsSupported                   []string    `json:"claims_supported"`
	RequestURIParameterSupported      bool        `json:"request_uri_parameter_supported"`
	TenantRegionScope                 interface{} `json:"tenant_region_scope"`
	CloudInstanceName                 string      `json:"cloud_instance_name"`
	CloudGraphHostName                string      `json:"cloud_graph_host_name"`
	MsgraphHost                       string      `json:"msgraph_host"`
}

type jwksKeys struct {
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

const (
	openIDConfigurationUri = "https://login.microsoftonline.com/common/v2.0/.well-known/openid-configuration"
)

func retrieveOpenIdConfiguration() (openIdConfiguration *openIDConfiguration, err error) {
	// create httpClient with timeout option
	var httpClient = &http.Client{Timeout: 10 * time.Second}
	res, err := httpClient.Get(openIDConfigurationUri)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &openIdConfiguration)
	if err != nil {
		return nil, err
	}

	return openIdConfiguration, err
}

func (o openIDConfiguration) RetrieveJwksKeys() (jwksKeys *jwksKeys, err error) {

	// create httpClient with timeout option
	var httpClient = &http.Client{Timeout: 10 * time.Second}
	res, err := httpClient.Get(o.JwksURI)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	mErr := json.Unmarshal(body, &jwksKeys)
	if err != nil {
		return nil, mErr
	}

	return jwksKeys, nil
}

func convertToPemBlock(key string) ([]byte, error) {

	beforeHeader := "-----BEGIN PUBLIC KEY-----\n"
	afterHeader := "\n-----END PUBLIC KEY-----"

	return []byte(beforeHeader + key + afterHeader), nil

}

func (j jwksKeys) getX5C(kid string) (string, error) {
	if kid == "" {
		return "", errors.New("token cannot be empty")
	}

	for _, key := range j.Keys {
		if key.Kid == kid {
			return key.X5C[0], nil
		}
	}

	return "", errors.New("no mathing kid found")

}

// GetClaimsFromAccessToken: only get the claims from the access token, it does not check if a token is valid
// because Microsoft says the following:
// Currently, access tokens issued by the v2.0 endpoint can be consumed only by Microsoft Services.
// Your apps shouldn't need to perform
// any validation or inspection of access tokens for any of the currently supported scenarios
// https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-v2-tokens
func GetClaimsFromAccessToken(accessToken string) (map[string]interface{}, error) {
	if accessToken == "" {
		return nil, errors.New("no access token provided")
	}

	token, _ := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return nil, errors.New("unable to extract claims")
	}

}

// ValidateIdToken validated the OpenID connect token
func ValidateIdToken(idTokenString string) (authenticated bool, claimMap map[string]interface{}, err error) {

	if idTokenString == "" {
		return authenticated, claimMap, errors.New("no id token provided")
	}

	token, err := jwt.Parse(idTokenString, func(token *jwt.Token) (interface{}, error) {

		openIDConfiguration, err := retrieveOpenIdConfiguration()
		if err != nil {
			return nil, err
		}

		jwksKeys, err := openIDConfiguration.RetrieveJwksKeys()
		if err != nil {
			return nil, err
		}

		x5c, err := jwksKeys.getX5C(token.Header["kid"].(string))

		pemBlock, err := convertToPemBlock(x5c)
		if err != nil {
			return nil, err
		}

		return jwt.ParseRSAPublicKeyFromPEM(pemBlock)

	})

	if err != nil {
		return false, nil, err
	}

	if !token.Valid {
		return token.Valid, nil, errors.New("token invalid")
	}

	claims := token.Claims.(jwt.MapClaims)
	return token.Valid, claims, nil

}
