package service

import (
	"context"
	"time"

	"github.com/wishrem/goligoli/danmu/model"
	"github.com/wishrem/goligoli/danmu/proto/pb"
	"github.com/yitter/idgenerator-go/idgen"
)

type DanmuService struct {
	pb.UnimplementedDanmuServiceServer
}

func (ds *DanmuService) Send(ctx context.Context, req *pb.SendReq) (*pb.SendResp, error) {
	dm := &model.Danmu{
		ID:      idgen.NextId(),
		UserID:  req.UserId,
		VideoID: req.VideoId,
		Content: req.Content,
		BeginAt: time.Unix(req.BeginAt, 0),
	}
	if err := dm.Create(); err != nil {
		return nil, err
	}

	return parseSendResp(dm), nil
}
func (ds *DanmuService) GetDanmus(ctx context.Context, req *pb.GetDanmusReq) (*pb.GetDanmusResp, error) {
	dm := &model.Danmu{
		VideoID: req.VideoId,
	}
	dms := make([]model.Danmu, 0)
	if err := dm.GetAllByVideoID(&dms); err != nil {
		return nil, err
	}

	return parseGetDanmusResp(&dms), nil
}
