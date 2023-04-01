package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wishrem/goligoli/api-gateway/service"
	danmu "github.com/wishrem/goligoli/danmu/proto/pb"
	"github.com/wishrem/goligoli/logger"
)

func SendDanmu(c *gin.Context) {
	vid := new(VideoID)
	if err := c.ShouldBindUri(vid); err != nil {
		logger.Log.Debugln(err)
		SendBadRequest(c)
		return
	}

	claims := ParseToken(c)
	if claims == nil {
		return
	}

	req := new(danmu.SendReq)
	if err := c.ShouldBind(req); err != nil {
		logger.Log.Debugln(err)
		SendBadRequest(c)
		return
	}
	req.UserId = claims.UserID
	req.VideoId = vid.ID

	resp, err := service.DanmuClient.Send(req)
	if err != nil {
		logger.Log.Debugln(err)
		SendErrResp(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetDanmus(c *gin.Context) {
	vid := new(VideoID)
	if err := c.ShouldBindUri(vid); err != nil {
		logger.Log.Debugln(err)
		SendBadRequest(c)
		return
	}

	claims := ParseToken(c)
	if claims == nil {
		return
	}

	req := new(danmu.GetDanmusReq)
	if err := c.ShouldBind(req); err != nil {
		logger.Log.Debugln(err)
		SendBadRequest(c)
		return
	}
	req.VideoId = vid.ID

	resp, err := service.DanmuClient.GetDanmus(req)
	if err != nil {
		logger.Log.Debugln(err)
		SendErrResp(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
