package main

import (
	"admin-api/common/config"
	_ "admin-api/docs"
	"admin-api/pkg/db"
	"admin-api/pkg/log"
	"admin-api/pkg/redis"
	"admin-api/router"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title 通用后台管理系统
// @version 1.0
// @description 后台管理系统API接口文档
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// 加载日志log
	logger := log.Log()
	// 设置启动模式
	gin.SetMode(config.Config.Server.Mode)
	// 初始化路由
	router := router.InitRouter()
	srv := &http.Server{
		Addr: config.Config.Server.Address,
		Handler: router,
	}
	// 启动服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Info("listen failed: %v \n", config.Config.Server.Address)
		}
		logger.Info("listen: %v \n", config.Config.Server.Address)
	}()
	quit := make(chan os.Signal) //监听消息
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Info("Server Shutdown:", err)
	}
	logger.Info("Server exiting")
}

// 初始化连接
func init() {
  	// mysql
  	db.SetUpDbLink()
	// redis
	redis.SetUpRedisDb()
}
