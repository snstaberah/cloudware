package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"inspur.com/cloudware/controller"
	"inspur.com/cloudware/controller/host"
	"inspur.com/cloudware/controller/user"
	_ "inspur.com/cloudware/docs"
	"inspur.com/cloudware/middleware"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	// r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	// r.Use(middleware.CurrentUser())
	// 开启swag路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", controller.Ping)

		// 使用 ?name=Sheldon传参数
		v1.GET("testlog", controller.TestLog)

		// 使用 body  json格式传参数
		v1.GET("testpost", controller.TestPost)

		// 使用 body  json格式传参数
		v1.GET("testpoststruct", controller.TestPost)

		// 用户登录
		v1.POST("user/register", user.UserRegister)

		// 用户登录
		v1.POST("user/login", user.UserLogin)

		// AddHost
		v1.POST("host", host.AddHost)

		// GetHostByID
		v1.GET("host", host.GetHostByID)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.JWTAuthRequired())
		{
			// User Routing
			auth.GET("user/me", user.UserMe)
			auth.DELETE("user/logout", user.UserLogout)
		}
	}
	return r
}
