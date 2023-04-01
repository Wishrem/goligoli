package service

import (
	"context"
	"fmt"
	"time"

	comment "github.com/wishrem/goligoli/comment/proto/pb"
	"github.com/wishrem/goligoli/pkg/conf"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var CommentClient *commentClient

type commentClient struct {
	conn *grpc.ClientConn
}

func SetupCommentClient() {
	etcdClient, err := clientv3.NewFromURL(conf.App.Etcd.URL)
	if err != nil {
		panic(err)
	}

	etcdResolver, err := resolver.NewBuilder(etcdClient)
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial(fmt.Sprintf("etcd:///%s", conf.App.CommentService.Name), grpc.WithResolvers(etcdResolver), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	CommentClient = new(commentClient)
	CommentClient.conn = conn
}

func (cc *commentClient) Comment(req *comment.CommentReq) (*comment.CommentResp, error) {
	c := comment.NewCommentServiceClient(cc.conn)
	req.SentAt = time.Now().Unix()
	return c.Comment(context.Background(), req)
}

func (cc *commentClient) Response(req *comment.ResponseReq) (*comment.ResponseResp, error) {
	c := comment.NewCommentServiceClient(cc.conn)
	req.SentAt = time.Now().Unix()
	return c.Response(context.Background(), req)
}

func (cc *commentClient) GetComment(req *comment.GetCommentReq) (*comment.GetCommentResp, error) {
	c := comment.NewCommentServiceClient(cc.conn)
	return c.GetComment(context.Background(), req)
}
