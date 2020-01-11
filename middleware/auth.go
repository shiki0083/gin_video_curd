package middleware

import (
	"fmt"
	"singo/model"
	"singo/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"singo/middleware/jwtauth"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			user, err := model.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// auth.GET("user", func(c *gin.Context) {
		// 	claims := c.MustGet("claims").(*jwtauth.CustomClaims)
		// 	fmt.Println(claims.ID)
		// 	c.String(http.StatusOK, claims.Name)
		// })
		jwtauth.JWTAuth()
		fmt.Println(jwtauth.JWTAuth())
		fmt.Println(*c)
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}
