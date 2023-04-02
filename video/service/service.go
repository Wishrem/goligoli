package service

import (
	"context"
	"fmt"

	"github.com/wishrem/goligoli/erp"
	"github.com/wishrem/goligoli/logger"
	"github.com/wishrem/goligoli/pkg/conf"
	"github.com/wishrem/goligoli/pkg/util"
	"github.com/wishrem/goligoli/video/model"
	"github.com/wishrem/goligoli/video/proto/pb"
	"github.com/yitter/idgenerator-go/idgen"
)

type VideoService struct {
	pb.UnimplementedVideoServiceServer
}

func (vs *VideoService) Upload(ctx context.Context, req *pb.UploadReq) (*pb.UploadResp, error) {
	ss, err := util.GenerateFileName(req.Video)
	if err != nil {
		logger.Log.Debugln(err)
		return nil, erp.Internal
	}
	ss = ss + ".mp4"
	filename := conf.App.Res.VideoDir + ss
	if err := util.WriteFile(filename, req.Video, 0666); err != nil {
		logger.Log.Debugln(err)
		return nil, erp.Internal
	}

	v := &model.Video{
		ID:          idgen.NextId(),
		UserID:      req.UserId,
		Title:       req.Title,
		Description: req.Description,
		Liked:       0,
		Shared:      0,
		VideoUrl:    "127.0.0.1:" + conf.App.Gateway.Port + "/goligoli/view/video/" + ss,
		Status: &model.Status{
			ID:     idgen.NextId(),
			Passed: false,
		},
	}
	if err := v.Create(); err != nil {
		logger.Log.Debugln(err)
		return nil, erp.Internal
	}

	return &pb.UploadResp{
		Video: generatePBVideo(v),
	}, nil
}

func (vs *VideoService) Like(ctx context.Context, req *pb.LikeReq) (*pb.LikeResp, error) {
	v := &model.Video{ID: req.VideoId}
	return &pb.LikeResp{}, v.Like()
}

func (vs *VideoService) Share(ctx context.Context, req *pb.ShareReq) (*pb.ShareResp, error) {
	ShareID := fmt.Sprintf("%v", req.UserId)
	v := &model.Video{ID: req.VideoId}
	if err := v.Share(); err != nil {
		logger.Log.Debugln(err)
		return nil, err
	}

	return &pb.ShareResp{
		ShareUrl: v.VideoUrl + "?share_id=" + ShareID,
	}, nil
}

func (vs *VideoService) Judge(ctx context.Context, req *pb.JudgeReq) (*pb.JudgeResp, error) {
	v := &model.Video{
		ID: req.VideoId,
		Status: &model.Status{
			Reason: req.Reason,
			Passed: req.Passed,
		},
	}
	if err := v.JudgeThenGet(); err != nil {
		logger.Log.Debugln(err)
		return nil, err
	}

	return &pb.JudgeResp{
		Video: generatePBVideo(v),
	}, nil
}

func (vs *VideoService) View(ctx context.Context, req *pb.ViewReq) (*pb.ViewResp, error) {
	v := &model.Video{VideoUrl: "127.0.0.1:" + conf.App.Gateway.Port + "/goligoli/view/video/" + req.Filename}
	if err := v.View(); err != nil {
		return nil, err
	}

	if !v.Status.Passed {
		return nil, erp.VideoNotFound
	}
	return &pb.ViewResp{}, nil
}

func (vs *VideoService) GetVideos(ctx context.Context, req *pb.GetVideosReq) (*pb.GetVideosResp, error) {
	s := &model.Search{
		Year:   req.Year,
		Shared: req.Shared,
		Liked:  req.Liked,
		Title:  req.Title,
	}

	videos := make([]model.Video, 0)
	if err := s.SearchVideos(&videos); err != nil {
		logger.Log.Debugln(err)
		return nil, erp.Internal
	}

	return parseGetVideosResp(&videos), nil
}
