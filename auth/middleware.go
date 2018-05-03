package auth

import (
	"net/http"
	"strings"
)

//AzureJWTAuthMiddleware: handler that takes an open-id id token from the Authorization Header
func AzureJWTAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		authorizationHeader := req.Header.Get("Authorization")
		if authorizationHeader != "" {

			bearerToken := strings.Split(authorizationHeader, " ")

			valid, claims, err := ValidateIdToken(bearerToken[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
			}

			if valid {
				w.Header().Add("TOKEN_NAME", claims["name"].(string))
				w.Header().Add("id_token", bearerToken[1])
				next.ServeHTTP(w, req)
			}
		}
	})
}
