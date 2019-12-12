package api //api 就是控制器用来接受请求 ========> MVC框架

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	service := service.CreateVideoService{} // 创建一个CreateVideoService()具体方法 但是在控制器中不要写业务逻辑！！！！  --- 做业务分发
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
