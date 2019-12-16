package service

import (
	"singo/model"
	"singo/serializer"
)

// ShowVideoService 视频展示
type ShowVideoService struct {
}

// Show 查看视频
func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	// db.First(&user, 10)  使用主键获取记录（仅适用于整数主键）
	error := model.DB.First(&video, id).Error
	if error != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: error.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video), //序列化器
	}

}
