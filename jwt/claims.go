package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GetClaims(c *gin.Context) jwt.MapClaims {
	if claims, ok := c.Get("claims"); ok {
		return claims.(jwt.MapClaims)
	}
	return nil
}
