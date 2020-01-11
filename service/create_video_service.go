package service

import (
	"singo/model"
	"singo/serializer"
)

// CreateVideoService 创建标签
type CreateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=5,max=30"`
	Info  string `form:"info" json:"info" binding:"required,min=0,max=200"`
}

// Create 创建标签  //接受一个请求 -->通过序列化器 --> 返回一个JSON
func (service *CreateVideoService) Create() serializer.Response { //====>给结构体定义方法 只好这么做
	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
	}

	error := model.DB.Create(&video).Error // 把标签存到数据库 并 捕获可能发生的错误 .Error 展示错误

	if error != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "标签保存失败",
			Error: error.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video), //序列化器
	}

}
