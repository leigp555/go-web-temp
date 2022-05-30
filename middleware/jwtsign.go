package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

//设置私钥
var mySigningKey = []byte("AllYourBase")

type MyCustomClaims struct {
	User string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成token
func GenerateToken(ctx *gin.Context) {
	claims := MyCustomClaims{
		"bar",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Second)), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",     //颁发者
			Subject:   "somebody", //主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(mySigningKey)
	if err != nil {
		ctx.JSON(500, map[string]string{
			"msg": "服务器异常,请稍后再试",
		})
	} else {
		ctx.JSON(200, map[string]string{
			"token": tokenStr,
		})
	}
	ctx.Abort()
}

func ParseToken(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	token, _ := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	claims, ok := token.Claims.(*MyCustomClaims)
	if ok && token.Valid {
		ctx.JSON(200, map[string]string{
			"user": claims.User,
		})
	} else {
		ctx.JSON(400, map[string]string{
			"msg": "未认证用户",
		})
	}
	ctx.Abort()
}
