package service

import (
	"singo/model"
	"singo/serializer"
)

// ListVideoService 视频展示
type ListVideoService struct {
}

// List 视频列表
func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video
	// ///// Get all matched records
	// ///// db.Where("name = ?", "jinzhu").Find(&users)
	error := model.DB.Find(&videos).Error
	if error != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "数据库链接错误",
			Error: error.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideos(videos), //序列化器
	}

}
