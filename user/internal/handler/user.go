package handler

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/wishrem/goligoli/pkg/e"
	"github.com/wishrem/goligoli/pkg/jwt"
	"github.com/wishrem/goligoli/user/internal/model"
	"github.com/wishrem/goligoli/user/proto/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func getSHA256String(s string) string {
	msg := []byte(s)
	hash := sha256.New()
	hash.Write(msg)
	bytes := hash.Sum(nil)
	return hex.EncodeToString(bytes)
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

	if u.Password != getSHA256String(req.Password) {
		return nil, status.Error(codes.Canceled, e.USER_WRONG_PASSWORD)
	}

	token, err := jwt.Generate(u.ID, []string{"user"})
	if err != nil {
		return nil, status.Error(codes.Internal, e.INTERNAL)
	}
	return &pb.LoginResp{
		User: &pb.User{
			Name:        u.Name,
			Email:       u.Email,
			Description: u.Description,
			PhotoUrl:    u.PhotoUrl,
		},
		Token: token,
	}, nil
}

func (us *UserServer) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	u := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: getSHA256String(req.Password),
	}
	if err := u.Create(ctx); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, status.Error(codes.AlreadyExists, e.USER_HAS_EXISTED)
		}
		return nil, status.Error(codes.Internal, e.INTERNAL)
	}

	token, err := jwt.Generate(u.ID, []string{"user"})
	if err != nil {
		return nil, status.Error(codes.Internal, e.INTERNAL)
	}
	return &pb.RegisterResp{
		User: &pb.User{
			Name:        u.Name,
			Email:       u.Email,
			Description: u.Description,
			PhotoUrl:    u.PhotoUrl,
		},
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
		User: &pb.User{
			Name:        u.Name,
			Email:       u.Email,
			Description: u.Description,
			PhotoUrl:    u.PhotoUrl,
		},
	}, nil
}
