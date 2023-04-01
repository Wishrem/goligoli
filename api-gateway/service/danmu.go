package service

import (
	"context"
	"fmt"

	danmu "github.com/wishrem/goligoli/danmu/proto/pb"
	"github.com/wishrem/goligoli/pkg/conf"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var DanmuClient *danmuClient

type danmuClient struct {
	conn *grpc.ClientConn
}

func SetupDanmuClient() {
	etcdClient, err := clientv3.NewFromURL(conf.App.Etcd.URL)
	if err != nil {
		panic(err)
	}

	etcdResolver, err := resolver.NewBuilder(etcdClient)
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial(fmt.Sprintf("etcd:///%s", conf.App.DanmuService.Name), grpc.WithResolvers(etcdResolver), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	DanmuClient = new(danmuClient)
	DanmuClient.conn = conn
}

func (dc *danmuClient) Send(req *danmu.SendReq) (*danmu.SendResp, error) {
	d := danmu.NewDanmuServiceClient(dc.conn)
	return d.Send(context.Background(), req)
}

func (dc *danmuClient) GetDanmus(req *danmu.GetDanmusReq) (*danmu.GetDanmusResp, error) {
	d := danmu.NewDanmuServiceClient(dc.conn)
	return d.GetDanmus(context.Background(), req)
}
