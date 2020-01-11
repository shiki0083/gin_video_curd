package server

import (
	"os"
	"singo/api"
	"singo/middleware"

	"github.com/gin-gonic/gin"

	"singo/middleware/jwtauth"
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

		// 用户注册
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("/")
		auth.Use(jwtauth.JWTAuth())
		{
			// User Routing
			// auth.GET("user", func(c *gin.Context) {
			// 	claims := c.MustGet("claims").(*jwtauth.CustomClaims)
			// 	fmt.Println(claims.ID)
			// 	c.String(http.StatusOK, claims.Name)
			// })
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
			auth.POST("createVideo", api.CreateVideo)
		}

		v1.POST("showVideo", api.ShowVideo) //URL 参数通过 DefaultQuery 或 Query 方法获取
		v1.GET("videoList", api.ListVideo)
		v1.POST("videoUpate", api.UpdateVideo) //URL 参数通过 DefaultQuery 或 Query 方法获取
		v1.POST("videoDelete", api.DeleteVideo)
	}
	return r
}
