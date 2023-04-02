package service

import (
	"github.com/wishrem/goligoli/video/model"
	"github.com/wishrem/goligoli/video/proto/pb"
)

func generatePBVideo(v *model.Video) *pb.Video {
	video := new(pb.Video)
	video.Description = v.Description
	video.Liked = v.Liked
	video.Passed = v.Status.Passed
	video.Reason = v.Status.Reason
	video.Shared = v.Shared
	video.Title = v.Title
	video.UserId = v.UserID
	video.VideoId = v.ID
	video.VideoUrl = v.VideoUrl
	return video
}

func parseGetVideosResp(vs *[]model.Video) *pb.GetVideosResp {
	videos := make([]*pb.Video, 0)
	for _, v := range *vs {
		if v.Status != nil && v.Status.Passed {
			videos = append(videos, generatePBVideo(&v))
		}
	}
	return &pb.GetVideosResp{Videos: videos}
}
