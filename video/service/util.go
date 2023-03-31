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
