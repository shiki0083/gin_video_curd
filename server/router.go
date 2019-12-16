package server

import (
	"github.com/gin-gonic/gin"
	"os"
	"singo/api"
	"singo/middleware"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}
		v1.POST("video", api.CreateVideo)
		v1.GET("video", api.ShowVideo) //URL 参数通过 DefaultQuery 或 Query 方法获取
		v1.GET("videoList", api.ListVideo)
		v1.GET("videoUpate", api.UpdateVideo) //URL 参数通过 DefaultQuery 或 Query 方法获取
		v1.GET("videoDelete", api.DeleteVideo)
	}
	return r
}
