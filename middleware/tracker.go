package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"blog/zlog"
)

func Tracker() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		zlog.Info(fmt.Sprintf("RequestURI:%s", c.Request.RequestURI))
		c.Next()

		// 响应时间
		log.Println("Time:", time.Since(t).String())
	}
}
