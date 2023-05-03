package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")

}
