package serializer

import "singo/model"

// Video 用户序列化器
type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	CreatedAt int64  `json:"created_at"`
}

// BuildVideo 序列化用户                MVC 模式下 这个部分本来是后端展示的代码 ---- 前后端分离的情况下 这就就是
func BuildVideo(item model.Video) Video {
	return Video{
		ID:        item.ID, // model模板内自带的
		Title:     item.Title,
		Info:      item.Info,
		CreatedAt: item.CreatedAt.Unix(), // model模板内自带的
	}
}
