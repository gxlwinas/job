package middleware

import (
	"go_pro/common"
	"go_pro/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		// tokenAnyType, ok := c.Get("token")

		// if !ok {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未获取到token"})
		// 	c.Abort()
		// 	return
		// }

		// tokenString := tokenAnyType.(string)

		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			c.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			c.Abort()
			return
		}

		userID := claims.UserID
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userID)

		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
