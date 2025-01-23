package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// RoleMiddleware enforces role-based access control for Gin-Gonic
func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract user role from context (set during authentication)
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized: Role not found"})
			return
		}

		// Check if role is allowed
		roleStr, ok := role.(string)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized: Invalid role type"})
			return
		}

		for _, allowedRole := range allowedRoles {
			if strings.EqualFold(roleStr, allowedRole) {
				c.Next() // Role is allowed, proceed to the next handler
				return
			}
		}

		// Role not allowed
		c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden: Access denied"})
	}
}
