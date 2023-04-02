package handler

import (
	"context"

	"errors"
	"fmt"
	"time"

	"github.com/wishrem/goligoli/pkg/conf"
	"github.com/wishrem/goligoli/pkg/e"
	"github.com/wishrem/goligoli/pkg/util"
	"github.com/wishrem/goligoli/pkg/util/jwt"
	"github.com/wishrem/goligoli/user/model"
	"github.com/wishrem/goligoli/user/proto/pb"
	"github.com/wishrem/goligoli/user/role"
	"github.com/yitter/idgenerator-go/idgen"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (us *UserServer) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	u := &model.User{
		Email: req.Email,
	}
	if err := u.Get(ctx); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, e.USER_NOT_FOUND)
		}
		return nil, status.Error(codes.Internal, e.INTERNAL)
	}

	s, err := getSHA256String(req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, e.INTERNAL)
	}

	if u.Password != s {
		return nil, status.Error(codes.Canceled, e.USER_WRONG_PASSWORD)
	}

	token, err := jwt.Generate(u.ID, role.GetRole(u.Roles))
	if err != nil {
		return nil, status.Error(codes.Internal, e.INTERNAL)
	}
	return &pb.LoginResp{
		User:  parseUser(u),
		Token: token,
	}, nil
}

func (us *UserServer) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	s, err := getSHA256String(req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, e.INTERNAL)
	}

	u := &model.User{
		ID:       idgen.NextId(),
		Name:     req.Name,
		Email:    req.Email,
		PhotoUrl: "127.0.0.1:" + conf.App.Gateway.Port + "/goligoli/view/photo/default.jpg",
		Password: s,
		Roles: []*model.Role{
			{
				ID:   idgen.NextId(),
				Type: "user",
			},
		},
		Ban: &model.Ban{
			ID:      idgen.NextId(),
			Reason:  "",
			BanAt:   time.Unix(0, 0),
			UnbanAt: time.Unix(0, 0),
		},
	}
	if err := u.Create(ctx); err != nil {
		return nil, status.Error(codes.Internal, e.INTERNAL)
	}

	token, err := jwt.Generate(u.ID, role.User)
	if err != nil {
		return nil, status.Error(codes.Internal, e.INTERNAL)
	}
	return &pb.RegisterResp{
		User:  parseUser(u),
		Token: token,
	}, nil
}

func (us *UserServer) GetInfo(ctx context.Context, req *pb.GetInfoReq) (*pb.GetInfoResp, error) {
	u := &model.User{
		ID: req.Id,
	}
	if err := u.Get(ctx); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, e.USER_NOT_FOUND)
		}
		return nil, status.Error(codes.Internal, e.INTERNAL)
	}

	return &pb.GetInfoResp{
		User: parseUser(u),
	}, nil
}

func (us *UserServer) Ban(ctx context.Context, req *pb.BanReq) (*pb.BanResp, error) {
	u := &model.User{
		ID: req.UserId,
	}

	u.Ban = &model.Ban{
		ID:      idgen.NextId(),
		Reason:  req.Ban.Reason,
		BanAt:   time.Unix(req.Ban.BanAt, 0),
		UnbanAt: time.Unix(req.Ban.UnbanAt, 0),
	}

	if err := u.UpdateBan(ctx); err != nil {
		return nil, status.Error(codes.Internal, e.INTERNAL)
	}

	return &pb.BanResp{
		User: parseUser(u),
	}, nil
}

func (us *UserServer) ModifyInfo(ctx context.Context, req *pb.ModifyInfoReq) (*pb.ModifyInfoResp, error) {
	u := &model.User{
		ID: req.Id,
	}

	if len(req.Photo) != 0 {
		s, err := util.GenerateFileName(req.Photo)
		if err != nil {
			return nil, status.Error(codes.Internal, e.INTERNAL)
		}
		s = s + "." + req.PhotoType
		filename := conf.App.Res.PhotoDir + s
		if err := util.WriteFile(filename, req.Photo, 0666); err != nil {
			return nil, status.Error(codes.Internal, e.INTERNAL)
		}

		u.PhotoUrl = "127.0.0.1:" + conf.App.Gateway.Port + "/goligoli/view/photo/" + s
	}

	if req.Description != "" {
		u.Description = req.Description
	}

	fmt.Println(u)
	if err := u.UpdateInfo(ctx); err != nil {
		return nil, status.Error(codes.Internal, e.INTERNAL)
	}

	return &pb.ModifyInfoResp{
		User: parseUser(u),
	}, nil
}
