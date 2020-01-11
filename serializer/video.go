/*
 * @Author: mikey.zhaopeng
 * @Date: 2019-12-13 14:53:03
 * @Last Modified by: mikey.zhaopeng
 * @Last Modified time: 2019-12-13 15:50:44
 */
package serializer

import "singo/model"

// Video 用户序列化器
type Video struct {
	// ID uint `json:"id"`
	// Title       string `json:"title"`
	// Info        string `json:"info"`
	// CreatedAt   int64  `json:"created_at"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// BuildVideo 序列化视频                MVC 模式下 这个部分本来是后端展示的代码 ---- 前后端分离的情况下 这就就是
//               进                出
func BuildVideo(item model.Video) Video {
	return Video{
		// ID: item.ID, // model模板内自带的
		// Title:     item.Title,
		// Info:      item.Info,
		// CreatedAt: item.CreatedAt.Unix(), // model模板内自带的
		Name:        item.Name,
		Description: item.Description,
	}
}

// BuildVideos 序列化视频列表
func BuildVideos(items []model.Video) []Video {
	var videos []Video //指的就是序列化好的JSON数组
	for _, item := range items {
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos //指的就是序列化好的JSON数组 --> 导出了
}
