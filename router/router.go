package router

import (
	"admin-api/api/controller"
	"admin-api/common/config"
	"admin-api/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {


	router := gin.New()
	// 宕机时自动恢复
	router.Use(gin.Recovery())
	// 跨域中间件
	router.Use(middleware.Cors())
	// 图片访问路径静态文件夹可直接访问
	router.StaticFS(config.Config.ImageSettings.UploadDir,
	http.Dir(config.Config.ImageSettings.UploadDir))
	// 日志中间件
	router.Use(middleware.Logger())
	// 注册路由
	register(router)
	return router
}

// register 路由接口
func register(router *gin.Engine)  {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/api/captcha", controller.Captcha)
	router.POST("/api/login", controller.Login)
}
