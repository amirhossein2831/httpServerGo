package Middleware

import (
	"errors"
	"github.com/amirhossein2831/httpServerGo/src/Auth"
	"github.com/amirhossein2831/httpServerGo/src/http/Response"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			Response.NewJson().
				SetSuccess(false).
				SetStatusCode(http.StatusUnauthorized).
				SetData(errors.New("unauthorized")).
				Log().Send(w)
			return
		}
		tokenString = tokenString[len("Bearer "):]

		err := Auth.VerifyToken(tokenString)
		if err != nil {
			Response.NewJson().
				SetSuccess(false).
				SetStatusCode(http.StatusBadRequest).
				SetData(errors.New("invalid Token")).
				Log().Send(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
