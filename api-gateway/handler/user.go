package handler

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wishrem/goligoli/api-gateway/service"
	"github.com/wishrem/goligoli/pkg/conf"
	"github.com/wishrem/goligoli/pkg/e"
	"github.com/wishrem/goligoli/pkg/util/jwt"
	user "github.com/wishrem/goligoli/user/proto/pb"
)

func LoginUser(c *gin.Context) {
	req := new(user.LoginReq)
	if err := c.ShouldBind(&req); err != nil {
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

	if err := c.ShouldBind(&req); err != nil {
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

	token := new(Token)
	if err := c.ShouldBind(&token); err != nil {
		SendBadRequest(c)
		return
	}

	claims, err := jwt.Parse(token.SS)
	if err != nil {
		SendError(c, err)
		return
	}

	req := new(user.GetInfoReq)
	req.Id = claims.UserID
	resp, err := service.UserClient.GetInfo(req)
	if err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func BanUser(c *gin.Context) {
	token := new(Token)
	if err := c.ShouldBind(&token); err != nil {
		SendBadRequest(c)
		return
	}

	claims, err := jwt.Parse(token.SS)
	if err != nil {
		SendError(c, err)
		return
	}

	req := new(user.BanReq)
	if err := c.ShouldBind(&req); err != nil {
		SendBadRequest(c)
		return
	}

	req.AdminId = claims.UserID

	resp, err := service.UserClient.Ban(req)
	if err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func ModifyUserInfo(c *gin.Context) {
	token := new(Token)
	if err := c.ShouldBind(&token); err != nil {
		SendBadRequest(c)
		return
	}

	claims, err := jwt.Parse(token.SS)
	if err != nil {
		SendError(c, err)
		return
	}

	req := new(user.ModifyInfoReq)
	if err := c.Bind(req); err != nil {
		SendBadRequest(c)
		return
	}
	req.Id = claims.UserID

	file, err := c.FormFile("photo")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		SendError(c, err)
		return
	}

	if file == nil && req.Description == "" {
		SendBadRequest(c)
		return
	}

	if file != nil {
		data, err := file.Open()
		if err != nil {
			SendError(c, err)
			return
		}
		defer data.Close()

		bytes, err := io.ReadAll(data)
		if err != nil {
			SendError(c, err)
			return
		}

		req.Photo = bytes
		s := strings.Split(file.Filename, ".")
		ss := strings.ToLower(s[len(s)-1])
		if ss != "jpg" && ss != "png" {
			SendError(c, errors.New(e.USER_INVALID_PHOTO_TYPE))
			return
		}
		req.PhotoType = ss
	}

	resp, err := service.UserClient.ModifyInfo(req)
	if err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func ViewPhoto(c *gin.Context) {
	file := new(File)
	if err := c.ShouldBindUri(&file); err != nil {
		SendBadRequest(c)
		return
	}

	filename := conf.App.Res.PhotoDir + file.FileName
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusFound, "/goligoli/404")
		return
	}
	f.Close()

	c.File(filename)
}
