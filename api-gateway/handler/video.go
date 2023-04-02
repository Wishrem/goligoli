package handler

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wishrem/goligoli/api-gateway/service"
	"github.com/wishrem/goligoli/erp"
	"github.com/wishrem/goligoli/logger"
	"github.com/wishrem/goligoli/pkg/conf"
	video "github.com/wishrem/goligoli/video/proto/pb"
)

func UploadVideo(c *gin.Context) {
	req := new(video.UploadReq)
	if err := c.ShouldBind(&req); err != nil {
		logger.Log.Debugln(err)
		SendBadRequest(c)
	}

	claims := ParseToken(c)
	if claims == nil {
		return
	}
	req.UserId = claims.UserID

	file, err := c.FormFile("video")
	if err != nil {
		logger.Log.Debugln(err)
		SendBadRequest(c)
		return
	}

	tmp := strings.Split(file.Filename, ".")
	t := strings.ToLower(tmp[len(tmp)-1])
	if t != "mp4" {
		SendErrResp(c, erp.New(erp.BAD_REQUEST, "video file should be mp4 type"))
		return
	}

	resp, err := service.VideoClient.Upload(req, file)
	if err != nil {
		logger.Log.Debugln(err)
		SendErrResp(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func ShareVideo(c *gin.Context) {
	req := new(video.ShareReq)
	if err := c.ShouldBind(&req); err != nil {
		logger.Log.Debugln(err)
		SendBadRequest(c)
		return
	}

	claims := ParseToken(c)
	if claims == nil {
		return
	}
	req.UserId = claims.UserID

	resp, err := service.VideoClient.Share(req)
	if err != nil {
		logger.Log.Debugln(err)
		SendErrResp(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func LikeVideo(c *gin.Context) {
	req := new(video.LikeReq)
	if err := c.ShouldBind(&req); err != nil {
		logger.Log.Debugln(err)
		SendBadRequest(c)
		return
	}

	claims := ParseToken(c)
	if claims == nil {
		return
	}
	req.UserId = claims.UserID

	_, err := service.VideoClient.Like(req)
	if err != nil {
		logger.Log.Debugln(err)
		SendErrResp(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func Judge(c *gin.Context) {
	req := new(video.JudgeReq)
	if err := c.ShouldBind(&req); err != nil {
		logger.Log.Debugln(err)
		SendBadRequest(c)
		return
	}

	claims := ParseToken(c)
	if claims == nil {
		return
	}
	if !claims.Roles.IsAdmin() {
		SendErrResp(c, erp.Forbidden)
		return
	}
	req.AdminId = claims.UserID

	resp, err := service.VideoClient.Judge(req)
	if err != nil {
		logger.Log.Debugln(err)
		SendErrResp(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func ViewVideo(c *gin.Context) {
	req := new(video.ViewReq)
	file := new(File)
	if err := c.ShouldBindUri(&file); err != nil {
		logger.Log.Debugln(err)
		SendBadRequest(c)
		return
	}

	req.Filename = file.FileName
	_, err := service.VideoClient.View(req)
	if err != nil {
		logger.Log.Debugln(err)
		SendErrResp(c, err)
		return
	}

	fileName := conf.App.Res.VideoDir + file.FileName
	f, err := os.Open(fileName)
	if err != nil {
		logger.Log.Debugln(err)
		c.Redirect(http.StatusFound, "/goligoli/404")
		return
	}
	f.Close()

	c.File(fileName)
}

func GetVideos(c *gin.Context) {
	req := new(video.GetVideosReq)
	if err := c.ShouldBind(&req); err != nil {
		logger.Log.Debugln(err)
		fmt.Println(err)
		SendBadRequest(c)
		return
	}

	claims := ParseToken(c)
	if claims == nil {
		return
	}

	if !HasVideoSearchingOpt(c, req) {
		return
	}

	resp, err := service.VideoClient.GetVideos(req)
	if err != nil {
		SendErrResp(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
