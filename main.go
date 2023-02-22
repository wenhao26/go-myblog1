package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gin-gonic/gin"

	"blog/config"
	"blog/models"
	"blog/router"
	"blog/zlog"
)

func main() {
	// 初始化日志
	zlog.InitLogger()
	logger := zlog.GetLogger()
	defer logger.Sync()

	// 初始化MySQL
	models.InitDb()

	// 启动服务
	zlog.Info("服务启动中...")
	err := start()
	if err != nil {
		zlog.Panic(fmt.Sprintf("启动服务失败：", err))
	}
}

func start() error {
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(config.HttpPort),
		Handler: getEngine(),
	}
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctxFunc context.CancelFunc) {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGINT)
		for {
			select {
			case <-ch:
				ctxFunc()
				return
			}
		}
	}(cancel)
	go func() {
		<-ctx.Done()
		if err := server.Shutdown(context.Background()); err != nil {
			panic(fmt.Sprintf("停止服务失败：", err.Error()))
		}
	}()

	zlog.Info("服务启动成功！")

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		zlog.Debug("服务已正常关闭！")
		return nil
	}
	return err
}

func getEngine() *gin.Engine {
	gin.SetMode(func() string {
		if config.IsDev() {
			return gin.DebugMode
		}
		return gin.ReleaseMode
	}())

	r := gin.New()
	r.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "服务器内部错误",
		})
	}))

	// 初始化路由
	zlog.Info("初始化路由配置...")
	router.InitRouter(r)

	return r
}
