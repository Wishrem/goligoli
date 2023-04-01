package service

import (
	"github.com/wishrem/goligoli/danmu/model"
	"github.com/wishrem/goligoli/danmu/proto/pb"
)

func parseDanmu(dm *model.Danmu) *pb.Danmu {
	return &pb.Danmu{
		Id:      dm.ID,
		UserId:  dm.UserID,
		VideoId: dm.VideoID,
		Content: dm.Content,
		BeginAt: dm.BeginAt.Unix(),
	}
}

func parseSendResp(dm *model.Danmu) *pb.SendResp {
	return &pb.SendResp{
		Danmu: parseDanmu(dm),
	}
}

func parseGetDanmusResp(dms *[]model.Danmu) *pb.GetDanmusResp {
	resp := make([]*pb.Danmu, 0)
	for _, i := range *dms {
		resp = append(resp, parseDanmu(&i))
	}
	return &pb.GetDanmusResp{
		Danmus: resp,
	}
}
