package Middleware

import (
	"github.com/amirhossein2831/httpServerGo/src/Auth"
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"net/http"
	"strings"
)

func PermissionMiddleware(permission string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			hasPermission := checkPermission(r, permission)
			if hasPermission {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Forbidden", http.StatusForbidden)
			}
		})
	}
}

func checkPermission(r *http.Request, permission string) bool {
	var user model.User
	authHeader := r.Header.Get("Authorization")
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	claims, err := Auth.RetainClaim(tokenString)
	if err != nil {
		return false

	}

	err = DB.GetInstance().GetDb().Where("email = ?", claims["email"]).Preload("Roles").Preload("Roles.Permissions").First(&user).Error
	if err != nil {
		return false
	}

	for _, role := range user.Roles {
		for _, rolePermission := range role.Permissions {
			if rolePermission.Name == permission {
				return true
			}
		}
	}

	return false
}
