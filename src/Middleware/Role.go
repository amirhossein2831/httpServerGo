package Middleware

import (
	"github.com/amirhossein2831/httpServerGo/src/Auth"
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"net/http"
	"strings"
)

func RoleMiddleware(role string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			hasRole := checkRole(r, role)
			if hasRole {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Forbidden", http.StatusForbidden)
			}
		})
	}
}

func checkRole(r *http.Request, role string) bool {
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

	for _, sr := range user.Roles {
		if sr.Name == role {
			return true
		}
	}

	return false
}
