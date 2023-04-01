package service

import (
	"github.com/wishrem/goligoli/comment/model"
	"github.com/wishrem/goligoli/comment/proto/pb"
)

func parseCommentResp(cm *model.Comment) *pb.CommentResp {
	return &pb.CommentResp{
		Comment: &pb.Resp{
			Id:      cm.ID,
			Content: cm.Content,
			SentAt:  cm.SentAt.Unix(),
		},
	}
}

func parseResponseResp(rp *model.Response) *pb.ResponseResp {
	return &pb.ResponseResp{
		Response: &pb.Resp{
			Id:      rp.ID,
			Content: rp.Content,
			SentAt:  rp.SentAt.Unix(),
		},
		CommentId: rp.CommentID,
	}
}

func parseGetCommentResp(cm *model.Comment) *pb.GetCommentResp {
	rps := make([]*pb.Resp, 0)
	for _, i := range cm.Responds {
		rps = append(rps, parseResponseResp(&i).Response)
	}
	return &pb.GetCommentResp{
		Comment:  parseCommentResp(cm).Comment,
		Responds: rps,
	}
}
