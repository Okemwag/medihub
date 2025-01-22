package middleware

import (
	"net/http"
	"strings"
)

// RoleMiddleware enforces role-based access control
func RoleMiddleware(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract user role from context (set during authentication)
			role, ok := r.Context().Value("role").(string)
			if !ok || role == "" {
				http.Error(w, "Unauthorized: Role not found", http.StatusUnauthorized)
				return
			}

			// Check if role is allowed
			for _, allowedRole := range allowedRoles {
				if strings.EqualFold(role, allowedRole) {
					next.ServeHTTP(w, r)
					return
				}
			}

			// Role not allowed
			http.Error(w, "Forbidden: Access denied", http.StatusForbidden)
		})
	}
}
