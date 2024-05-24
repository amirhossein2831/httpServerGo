package Middleware

import (
	"errors"
	"github.com/amirhossein2831/httpServerGo/src/Auth"
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/http/Response"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"net/http"
	"strings"
)

func CustomPermissionMiddleware(permissionMap map[string][]string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			method := r.Method
			permissions, exists := permissionMap[method]
			if !exists {
				next.ServeHTTP(w, r)
			}

			var user model.User
			authHeader := r.Header.Get("Authorization")
			tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

			claims, err := Auth.RetainClaim(tokenString)
			if err != nil {
				Response.NewJson().
					SetSuccess(false).
					SetStatusCode(http.StatusForbidden).
					SetData(errors.New("forbidden")).
					Log().Send(w)
				return
			}

			err = DB.GetInstance().GetDb().Where("email = ?", claims["email"]).Preload("Roles").Preload("Roles.Permissions").First(&user).Error
			if err != nil {
				Response.NewJson().
					SetSuccess(false).
					SetStatusCode(http.StatusForbidden).
					SetData(errors.New("forbidden")).
					Log().Send(w)
				return
			}

			hasPermission := model.HasPermission(user, permissions)
			if !hasPermission {
				Response.NewJson().
					SetSuccess(false).
					SetStatusCode(http.StatusForbidden).
					SetData(errors.New("forbidden")).
					Log().Send(w)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
