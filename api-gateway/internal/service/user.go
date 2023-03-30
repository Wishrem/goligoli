package service

import (
	"context"
	"fmt"

	"github.com/wishrem/goligoli/pkg/conf"
	user "github.com/wishrem/goligoli/user/proto/pb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var UserClient *userClient

type userClient struct {
	conn *grpc.ClientConn
}

func SetupUserClient() {
	etcdClient, err := clientv3.NewFromURL(conf.App.Etcd.URL)
	if err != nil {
		panic(err)
	}

	etcdResolver, err := resolver.NewBuilder(etcdClient)
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial(fmt.Sprintf("etcd:///%s", conf.App.UserService.Name), grpc.WithResolvers(etcdResolver), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	UserClient = new(userClient)
	UserClient.conn = conn
}

func (uc *userClient) Login(req *user.LoginReq) (*user.LoginResp, error) {
	c := user.NewUserServiceClient(uc.conn)
	return c.Login(context.Background(), req)
}

func (uc *userClient) Register(req *user.RegisterReq) (*user.RegisterResp, error) {
	c := user.NewUserServiceClient(uc.conn)
	return c.Register(context.Background(), req)
}

func (uc *userClient) GetInfo(req *user.GetInfoReq) (*user.GetInfoResp, error) {
	c := user.NewUserServiceClient(uc.conn)
	return c.GetInfo(context.Background(), req)
}
