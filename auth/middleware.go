package auth

import (
	"net/http"
	"strings"
)

// AzureJWTAuthMiddleware handler that take an JWT token from the Authorization Header  and
// validates it against Azure AD
func AzureJWTAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		authorizationHeader := req.Header.Get("Authorization")
		if authorizationHeader != "" {

			bearerToken := strings.Split(authorizationHeader, " ")

			valid, claims, err := ValidateToken(bearerToken[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
			}

			if valid {
				w.Header().Add("TOKEN_NAME", claims["name"].(string))
				next.ServeHTTP(w, req)
			}
		}
	})
}
