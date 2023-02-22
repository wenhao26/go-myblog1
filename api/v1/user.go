package v1

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"blog/api"
	"blog/middleware"
	"blog/models"
	"blog/utils"
	"blog/utils/errmsg"
)

type User struct {
	*api.Base
}

var UserAPI = &User{
	Base: api.BaseAPI,
}

// 注册账号
func (u *User) Register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if models.CheckUser(username) {
		u.Failed(ctx, errmsg.UserUsed, errmsg.GetErrMsg(errmsg.UserUsed))
		return
	}

	uid := models.CreateUser(&models.User{
		Username: username,
		Password: utils.GenPassword(password),
		Role:     0,
	})
	if uid == 0 {
		u.Failed(ctx, api.FailedCode, "注册失败")
		return
	}
	u.Success(ctx, "注册成功", map[string]interface{}{"uid": uid})
}

// 登录账号
func (u *User) Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	info, code := models.CheckLogin(username, password)
	if code == api.SuccessCode {
		// JWT
		j := middleware.NewJWT()
		claims := middleware.MyClaims{
			AuthInfo: info,
			StandardClaims: jwt.StandardClaims{
				NotBefore: time.Now().Unix() - 100,
				ExpiresAt: time.Now().Unix() + 3*24*3600,
				Issuer:    "wwh",
			},
		}
		token, err := j.CreateToken(claims)
		if err != nil {
			u.Failed(ctx, api.FailedCode, "登录失败，请稍后再试~")
			return
		}

		u.Success(ctx, "登录成功", map[string]interface{}{
			"info":  info,
			"token": token,
		})
		return
	}
	u.Failed(ctx, code, errmsg.GetErrMsg(code))
}

// 获取登录用户信息
func (u *User) UserInfo(ctx *gin.Context) {
	authInfo, ok := ctx.Get("auth-info")
	if !ok {
		u.Failed(ctx, api.FailedCode, "数据获取异常")
		return
	}
	u.Success(ctx, "success", authInfo)
}
