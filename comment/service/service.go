package service

import (
	"context"
	"time"

	"github.com/wishrem/goligoli/comment/model"
	"github.com/wishrem/goligoli/comment/proto/pb"
	"github.com/wishrem/goligoli/erp"
	"github.com/wishrem/goligoli/logger"
	"github.com/yitter/idgenerator-go/idgen"
)

type CommentService struct {
	pb.UnimplementedCommentServiceServer
}

func (cs *CommentService) Comment(ctx context.Context, req *pb.CommentReq) (*pb.CommentResp, error) {
	cm := &model.Comment{
		ID:      idgen.NextId(),
		Content: req.Content,
		SentAt:  time.Unix(req.SentAt, 0),
	}

	if err := cm.CreateComment(); err != nil {
		logger.Log.Debugln(err)
		return nil, erp.Internal
	}

	return parseCommentResp(cm), nil
}

func (cs *CommentService) Response(ctx context.Context, req *pb.ResponseReq) (*pb.ResponseResp, error) {
	rp := &model.Response{
		ID:        idgen.NextId(),
		CommentID: req.CommentId,
		Content:   req.Content,
		SentAt:    time.Unix(req.SentAt, 0),
	}

	if err := rp.CreateResponse(); err != nil {
		logger.Log.Debugln(err)
		return nil, erp.Internal
	}

	return parseResponseResp(rp), nil
}
func (cs *CommentService) GetComment(ctx context.Context, req *pb.GetCommentReq) (*pb.GetCommentResp, error) {
	cm := &model.Comment{
		ID: req.CommentId,
	}

	if err := cm.GetCommentAndResponds(); err != nil {
		logger.Log.Debugln(err)
		return nil, erp.Internal
	}

	return parseGetCommentResp(cm), nil
}
