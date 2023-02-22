package router

import (
	"time"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	limit2 "github.com/yangxikun/gin-limit-by-key"
	"golang.org/x/time/rate"

	"blog/api"
	v1 "blog/api/v1"
	"blog/middleware"
)

func InitRouter(r *gin.Engine) {
	// 设置允许最大访问数
	r.Use(limit.MaxAllowed(1000))
	// 根据客户端IP，限制访问速率
	r.Use(limit2.NewRateLimiter(func(c *gin.Context) string {
		return c.ClientIP()
	}, func(c *gin.Context) (*rate.Limiter, time.Duration) {
		// 限制每个IP，限制在10qps
		// 允许最多10个令牌并发，限制器活动时间为1小时
		return rate.NewLimiter(rate.Every(100*time.Millisecond), 10), time.Hour
	}, func(c *gin.Context) {
		c.AbortWithStatus(429)
	}))
	// 跟踪访问路由信息
	r.Use(middleware.Tracker())

	expire := time.Minute
	store := persistence.NewInMemoryStore(expire)

	// 默认路由
	r.GET("/", cache.CachePage(store, expire, api.BaseAPI.Default))

	group := r.Group("api/v1")
	{
		// 注册账号
		group.POST("register", v1.UserAPI.Register)
		// 登录账号
		group.POST("login", v1.UserAPI.Login)
	}

	authGroup := r.Group("api/v1").Use(middleware.JwtToken())
	{
		// 获取登录用户信息
		authGroup.GET("userinfo", cache.CachePage(store, expire, v1.UserAPI.UserInfo))
		// 获取文章列表
		authGroup.GET("articles", v1.ArticleAPI.Articles)
		// 获取文章信息
		authGroup.GET("article", v1.ArticleAPI.Article)

		// TODO...
	}

}
