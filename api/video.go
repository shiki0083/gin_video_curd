package api //api 就是控制器用来接受请求 ========> MVC框架

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// CreateVideo 视频投稿            ===================> 增
func CreateVideo(c *gin.Context) {
	service := service.CreateVideoService{} // 创建一个CreateVideoService()具体方法 但是在控制器中不要写业务逻辑！！！！  --- 做业务分发
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowVideo 视频详情接口			 ===================> 查
func ShowVideo(c *gin.Context) {
	service := service.ShowVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.DefaultQuery("id", "")) // ==============> 通过gin获取到的http请求内容都在 c 里，再传到Show()业务函数里
		c.JSON(200, res)                              //////////////////===== c.Param("id") api 参数通过Context的Param方法来获取
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListVideo 视频列表接口			 ===================> 查
func ListVideo(c *gin.Context) {
	service := service.ListVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateVideo 视频更新接口			 ===================> 改
func UpdateVideo(c *gin.Context) {
	service := service.UpdateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.DefaultQuery("id", "")) // ====>  URL 参数通过 DefaultQuery 或 Query 方法获取
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteVideo 视频更新接口			 ===================> 删
func DeleteVideo(c *gin.Context) {
	service := service.DeleteVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete(c.DefaultQuery("id", "")) // ====>  URL 参数通过 DefaultQuery 或 Query 方法获取
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
