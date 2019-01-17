package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/isayme/go-user/src/constant"
	"github.com/isayme/go-user/src/httperror"
	"github.com/isayme/go-user/src/jwt"
)

const tokenPrefix = "Bearer "

func AuthorizeRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := mustMetToken(c)

		claims, err := jwt.Verify(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		userID := claims.Get("uid")

		c.Set(constant.ClaimsKey, claims)
		c.Set(constant.UserIDKey, userID)
	}
}

func mustMetToken(c *gin.Context) string {
	value := c.GetHeader(constant.HTTPHeaderAuthorization)
	if value == "" {
		panic(httperror.AccessTokenRequired)
	}

	if !strings.HasPrefix(value, tokenPrefix) {
		panic(httperror.AccessTokenRequired)
	}

	return strings.TrimSpace(value[len(tokenPrefix):])
}
