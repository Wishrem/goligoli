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

		_goligoli.GET("/view/:type/:name", handler.View)
	}

	r.Run(":" + conf.App.Gateway.Port)
}
