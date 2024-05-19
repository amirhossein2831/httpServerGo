package Middleware

import (
	"github.com/amirhossein2831/httpServerGo/src/Auth"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenString = tokenString[len("Bearer "):]

		err := Auth.VerifyToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid Token", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
