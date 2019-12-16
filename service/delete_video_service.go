package service

import (
	"singo/model"
	"singo/serializer"
)

// DeleteVideoService 删除视频的服务
type DeleteVideoService struct {
}

// Delete 删除视频
func (service *DeleteVideoService) Delete(id string) serializer.Response { //====>给结构体定义方法 只好这么做
	video := model.Video{}
	// Delete an existing record 删除记录
	// db.Delete(&email)
	error := model.DB.First(&video, id).Error // 把视频存到数据库 并 捕获可能发生的错误 .Error 展示错误
	if error != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: error.Error(),
		}
	}
	error = model.DB.Delete(&video).Error // 把视频存到数据库 并 捕获可能发生的错误 .Error 展示错误
	if error != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "视频删除失败",
			Error: error.Error(),
		}
	}
	return serializer.Response{}
}
