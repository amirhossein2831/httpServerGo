package Middleware

import (
	"github.com/amirhossein2831/httpServerGo/src/Auth"
	"github.com/amirhossein2831/httpServerGo/src/DB"
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
			hasPermission := checkPermissions(r, permissions)
			if hasPermission {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Forbidden", http.StatusForbidden)
			}
		})
	}
}

func checkPermissions(r *http.Request, permissions []string) bool {
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
			for _, permission := range permissions {
				if rolePermission.Name == permission {
					return true
				}
			}
		}
	}

	return false
}
