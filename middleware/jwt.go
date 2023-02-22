package middleware

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"blog/api"
	"blog/config"
	"blog/utils/errmsg"
)

type JWT struct {
	JwtKey []byte
}

type MyClaims struct {
	//Username string `json:"username"`
	AuthInfo interface{}
	jwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		JwtKey: []byte(config.JwtKey),
	}
}

// 创建token
func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

// 解析token
func (j *JWT) ParseToken(verifyToken string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(verifyToken, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New(errmsg.GetErrMsg(errmsg.TokenMalformed))
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New(errmsg.GetErrMsg(errmsg.TokenExpired))
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New(errmsg.GetErrMsg(errmsg.TokenNotValidYet))
			} else {
				return nil, errors.New(errmsg.GetErrMsg(errmsg.TokenInvalid))
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New(errmsg.GetErrMsg(errmsg.TokenInvalid))
	}
	return nil, errors.New(errmsg.GetErrMsg(errmsg.TokenInvalid))
}

// JWT中间件
func JwtToken() gin.HandlerFunc {
	var code int
	var base api.Base

	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		token := strings.Split(authorization, " ")
		if authorization == "" || len(token) == 0 || len(token) != 2 || token[0] != "Bearer" {
			code = errmsg.TokenNotExist
			base.Failed(c, code, errmsg.GetErrMsg(code))
			c.Abort()
			return
		}

		j := NewJWT()

		// 解析token
		claims, err := j.ParseToken(token[1])
		if err != nil {
			if err == errors.New(errmsg.GetErrMsg(errmsg.TokenMalformed)) {
				base.Failed(c, errmsg.TokenMalformed, errmsg.GetErrMsg(errmsg.TokenMalformed))
			} else if err == errors.New(errmsg.GetErrMsg(errmsg.TokenExpired)) {
				base.Failed(c, errmsg.TokenExpired, errmsg.GetErrMsg(errmsg.TokenExpired))
			} else if err == errors.New(errmsg.GetErrMsg(errmsg.TokenNotValidYet)) {
				base.Failed(c, errmsg.TokenNotValidYet, errmsg.GetErrMsg(errmsg.TokenNotValidYet))
			} else {
				base.Failed(c, errmsg.TokenInvalid, errmsg.GetErrMsg(errmsg.TokenInvalid))
			}
			c.Abort()
			return
		}

		c.Set("auth-info", claims)
		c.Next()
	}
}
