package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wishrem/goligoli/api-gateway/handler"
	"github.com/wishrem/goligoli/pkg/conf"
)

func Setup() {
	r := gin.Default()

	_goligoli := r.Group("/goligoli")
	{
		_goligoli.GET("/user", handler.GetUserInfo)
		_goligoli.POST("/user", handler.ModifyUserInfo)
		_goligoli.POST("/user/login", handler.LoginUser)
		_goligoli.POST("/user/register", handler.RegisterUser)
		_goligoli.POST("/user/ban", handler.BanUser)

		// video
		_goligoli.POST("/video", handler.UploadVideo)
		_goligoli.GET("/video", handler.GetVideos)
		_goligoli.PUT("/video/status", handler.Judge)
		_goligoli.GET("/video/link", handler.ShareVideo)
		_goligoli.PUT("/video", handler.LikeVideo)

		// comment & response
		_goligoli.POST("/comment/:video_id", handler.CommentVideo)
		_goligoli.POST("/respond/:comment_id", handler.ResponseComment)
		_goligoli.GET("/comment/:comment_id", handler.GetComment)

		// danmu
		_goligoli.POST("/danmu/:video_id", handler.SendDanmu)
		_goligoli.GET("/danmu/:video_id", handler.GetDanmus)

		// view
		_goligoli.GET("/view/video/:name", handler.ViewVideo)
		_goligoli.GET("/view/photo/:name", handler.ViewPhoto)
	}

	r.Run(":" + conf.App.Gateway.Port)
}
