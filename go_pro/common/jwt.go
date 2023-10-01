package common

import (
	"go_pro/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserID uint
	jwt.StandardClaims
}
//发送token
func ReleaseToken(user model.User) (string, error) {
	//token的到期时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	//声明
	claims := &Claims{

		UserID: user.ID, //用户ID
		//负载 Payload部分
		StandardClaims: jwt.StandardClaims{

			ExpiresAt: expirationTime.Unix(), //失效于
			IssuedAt:  time.Now().Unix(), //发送于
			Issuer:    "127.0.0.1", //发布者
			Subject:   "user token",//主题
		},
	}
	//signature 签名部分
	
	//创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//获取完整的已签名令牌
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
//解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
