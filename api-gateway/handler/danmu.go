package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wishrem/goligoli/api-gateway/service"
	comment "github.com/wishrem/goligoli/comment/proto/pb"
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

	req := new(comment.CommentReq)
	if err := c.ShouldBind(req); err != nil {
		logger.Log.Debugln(err)
		SendBadRequest(c)
		return
	}
	req.VideoId = vid.ID

	resp, err := service.CommentClient.Comment(req)
	if err != nil {
		logger.Log.Debugln(err)
		SendErrResp(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
