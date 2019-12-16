package service

import (
	"singo/model"
	"singo/serializer"
)

// UpdateVideoService 视频投稿
type UpdateVideoService struct { //--------- 结构体
	Title string `form:"title" json:"title" binding:"required,min=5,max=30"`
	Info  string `form:"info" json:"info" binding:"required,min=0,max=200"`
}

// Update 更新视频
func (service *UpdateVideoService) Update(id string) serializer.Response {
	var video model.Video
	error := model.DB.First(&video, id).Error
	if error != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: error.Error(),
		}
	}

	video.Title = service.Title
	video.Info = service.Info
	error = model.DB.Save(&video).Error
	if error != nil {
		return serializer.Response{
			Code:  500002,
			Msg:   "视频修改保存失败",
			Error: error.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}

}
