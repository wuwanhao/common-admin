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

	// jwt鉴权接口
	jwt := router.Group("/api", middleware.AuthMiddleware())
	{
		jwt.POST("/api/post/add", controller.CreateSysPost)
		jwt.GET("/api/post/info", controller.GetSysPostById)
		jwt.PUT("/api/post/update", controller.UpdateSysPost)
		jwt.DELETE("/api/post/delete", controller.DeleteSysPostById)
		jwt.DELETE("/api/post/batch/delete", controller.BatchDeleteSysPost)
		jwt.PUT("/api/post/updateStatus", controller.UpdateSysPostStatus)
		jwt.GET("/api/post/list", controller.GetSysPostList)
		jwt.GET("/api/post/vo/list", controller.QuerySysPostVoList)
		jwt.POST("/api/dept/add", controller.CreateSysDept)
		jwt.GET("/api/dept/vo/list", controller.QuerySysDeptVoList)
		jwt.GET("/api/dept/info", controller.GetSysDeptById)
		jwt.PUT("/api/dept/update", controller.UpdateSysDept)
		jwt.DELETE("/api/dept/delete", controller.DeleteSysDeptById)
		jwt.GET("/api/dept/list", controller.GetSysDeptList)
		jwt.GET("/api/menu/vo/list", controller.QuerySysMenuVoList)
		jwt.POST("/api/menu/add", controller.CreateSysMenu)
		jwt.GET("/api/menu/info", controller.GetSysMenu)
		jwt.PUT("/api/menu/update", controller.UpdateSysMenu)
		jwt.DELETE("/api/menu/delete", controller.DeleteSysRoleMenu)
		jwt.GET("/api/menu/list", controller.GetSysMenuList)
		jwt.POST("/api/role/add", controller.CreateSysRole)
		jwt.PUT("/api/role/info", controller.GetSysRoleBYId)
		jwt.PUT("/api/role/update", controller.UpdateSysRole)
		jwt.DELETE("/api/role/delete", controller.DeleteSysRoleById)
		jwt.PUT("/api/role/updateStatus", controller.UpdateSysRoleStatus)
		jwt.GET("/api/role/list", controller.GetSysRoleList)
		jwt.GET("/api/role/vo/list", controller.QuerySysRoleVoList)
		jwt.GET("/api/role/vo/idList", controller.QueryRoleMenuIdList)
		jwt.PUT("/api/role/assignPermissions", controller.AssignPermissions)
		jwt.POST("/api/admin/add", controller.CreateSysAdmin)
		jwt.GET("/api/admin/info", controller.GetSysAdminInfo)
		jwt.PUT("/api/admin/update", controller.UpdateSysAdmin)
		jwt.DELETE("/api/admin/delete", controller.DeleteSysAdminById)
		jwt.PUT("/api/admin/updateStatus", controller.UpdateSysAdminStatus)
		jwt.PUT("/api/admin/updatePassword", controller.ResetSysAdminPassword)
		jwt.PUT("/api/admin/updatePersonal", controller.UpdatePersonal)
		jwt.PUT("/api/admin/updatePersonalPassword", controller.UpdatePersonalPassword)
		jwt.GET("/api/admin/list", controller.GetSysAdminList)
		jwt.POST("/upload", controller.Upload)
		jwt.GET("/sysLoginInfo/list", controller.GetSysLoginInfoList)
		jwt.DELETE("/sysLoginInfo/batch/delete", controller.BatchDeleteSysLoginInfo)
		jwt.DELETE("/sysLoginInfo/delete", controller.DeleteSysLoginInfoById)
		jwt.DELETE("/sysLoginInfo/clean", controller.CleanSysLoginInfo)
	}


}
