package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func JwtAccessTokenValidator(next http.Handler) http.Handler {
   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

   })
}


