package api

import (
	"fmt"
	"singo/serializer"
	"github.com/gin-gonic/gin"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	fmt.Println("connection succedssed",c)
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "成功",
	})
}