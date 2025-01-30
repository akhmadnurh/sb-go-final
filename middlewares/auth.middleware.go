package middlewares

import (
	"OneTix/utils"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

func RoleAuth(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		AuthToken := c.Request.Header.Get("Authorization")

		if AuthToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		token := strings.Split(AuthToken, " ")[1]

		user, err := utils.VerifyToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if !slices.Contains(roles, user.Role) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access Forbidden"})
			c.Abort()
			return
		}

		c.Set("user", *user)

		c.Next()
	}
}
