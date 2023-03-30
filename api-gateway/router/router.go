package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wishrem/goligoli/api-gateway/internal/handler"
)

func Setup() {
	r := gin.Default()

	_goligoli := r.Group("/goligoli")
	{
		_goligoli.GET("/user", handler.GetUserInfo)
		_goligoli.POST("/user/login", handler.LoginUser)
		_goligoli.POST("/user/register", handler.RegisterUser)
	}

	r.Run()
}
