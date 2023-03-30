package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wishrem/goligoli/api-gateway/internal/service"
	"github.com/wishrem/goligoli/pkg/jwt"
	user "github.com/wishrem/goligoli/user/proto/pb"
)

func LoginUser(c *gin.Context) {
	req := new(user.LoginReq)
	err := c.ShouldBind(&req)
	if err != nil {
		SendBadRequest(c)
		return
	}

	resp, err := service.UserClient.Login(req)
	if err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func RegisterUser(c *gin.Context) {
	req := new(user.RegisterReq)
	err := c.ShouldBind(&req)
	fmt.Println(req)
	if err != nil {
		SendBadRequest(c)
	}

	resp, err := service.UserClient.Register(req)
	if err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetUserInfo(c *gin.Context) {
	req := new(user.GetInfoReq)
	token := new(Token)
	err := c.ShouldBind(&token)
	if err != nil {
		SendBadRequest(c)
		return
	}

	claims, err := jwt.Parse(token.SS)
	if err != nil {
		SendError(c, err)
		return
	}
	req.Id = claims.UserID
	resp, err := service.UserClient.GetInfo(req)
	if err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
