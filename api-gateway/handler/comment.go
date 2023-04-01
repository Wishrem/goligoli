package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wishrem/goligoli/api-gateway/service"
	comment "github.com/wishrem/goligoli/comment/proto/pb"
	"github.com/wishrem/goligoli/logger"
)

func CommentVideo(c *gin.Context) {
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

func ResponseComment(c *gin.Context) {
	cid := new(CommentID)
	if err := c.ShouldBindUri(cid); err != nil {
		logger.Log.Debugln(err)
		SendBadRequest(c)
		return
	}

	claims := ParseToken(c)
	if claims == nil {
		return
	}

	req := new(comment.ResponseReq)
	if err := c.ShouldBind(req); err != nil {
		logger.Log.Debugln(err)
		SendBadRequest(c)
		return
	}
	req.CommentId = cid.ID

	resp, err := service.CommentClient.Response(req)
	if err != nil {
		logger.Log.Debugln(err)
		SendErrResp(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetComment(c *gin.Context) {
	cid := new(CommentID)
	if err := c.ShouldBindUri(cid); err != nil {
		logger.Log.Debugln(err)
		SendBadRequest(c)
		return
	}

	claims := ParseToken(c)
	if claims == nil {
		return
	}

	req := &comment.GetCommentReq{
		CommentId: cid.ID,
	}
	resp, err := service.CommentClient.GetComment(req)
	if err != nil {
		logger.Log.Debugln(err)
		SendErrResp(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
