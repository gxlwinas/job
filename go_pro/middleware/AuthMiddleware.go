package middleware

import (
	"go_pro/common"
	"go_pro/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)
//闭包
func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		// tokenAnyType, ok := c.Get("token")

		// if !ok {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未获取到token"})
		// 	c.Abort()
		// 	return
		// }

		// tokenString := tokenAnyType.(string)

		//从Header的Authorization中获取t签名
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			c.Abort()
			return
		}
		//取签名的有效位数。第7位开始
		tokenString = tokenString[7:]
		//解析签名
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			c.Abort()
			return
		}

		userID := claims.UserID
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userID) //查找userID

		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			c.Abort()
			return
		}
		//在上下文中设置key "user"
		c.Set("user", user)
		//执行下一个ginHandler
		c.Next()
	}
}
