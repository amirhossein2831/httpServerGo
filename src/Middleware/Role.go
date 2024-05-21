package Middleware

import (
	"github.com/amirhossein2831/httpServerGo/src/Auth"
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"net/http"
	"strings"
)

func RoleMiddleware(roles []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var user model.User

			authHeader := r.Header.Get("Authorization")
			tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

			claims, err := Auth.RetainClaim(tokenString)
			if err != nil {
				http.Error(w, "Forbidden", http.StatusForbidden)
			}

			err = DB.GetInstance().GetDb().Where("email = ?", claims["email"]).Preload("Roles").Preload("Roles.Permissions").First(&user).Error
			if err != nil {
				http.Error(w, "Forbidden", http.StatusForbidden)
			}

			hasRole := model.HasRole(user, roles)
			if hasRole {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Forbidden", http.StatusForbidden)
			}
		})
	}
}
