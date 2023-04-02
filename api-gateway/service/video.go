package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"github.com/wishrem/goligoli/erp"
	"github.com/wishrem/goligoli/pkg/conf"
	video "github.com/wishrem/goligoli/video/proto/pb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var VideoClient *videoClient

type videoClient struct {
	conn *grpc.ClientConn
}

func SetupVideoClient() {
	etcdClient, err := clientv3.NewFromURL(conf.App.Etcd.URL)
	if err != nil {
		panic(err)
	}

	etcdResolver, err := resolver.NewBuilder(etcdClient)
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial(fmt.Sprintf("etcd:///%s", conf.App.VideoService.Name), grpc.WithResolvers(etcdResolver), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	VideoClient = new(videoClient)
	VideoClient.conn = conn
}

func (vc *videoClient) Upload(req *video.UploadReq, file *multipart.FileHeader) (*video.UploadResp, error) {
	v := video.NewVideoServiceClient(vc.conn)

	f, err := file.Open()
	if err != nil {
		Log.Debugln(err)
		return nil, err
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		Log.Debugln(err)
		return nil, err
	}
	req.Video = bytes

	req.Year = int64(time.Now().Year())

	return v.Upload(context.Background(), req)
}

func (vc *videoClient) Share(req *video.ShareReq) (*video.ShareResp, error) {
	v := video.NewVideoServiceClient(vc.conn)
	return v.Share(context.Background(), req)
}

func (vc *videoClient) Like(req *video.LikeReq) (*video.LikeResp, error) {
	v := video.NewVideoServiceClient(vc.conn)
	return v.Like(context.Background(), req)
}

func (vc *videoClient) View(req *video.ViewReq) (*video.ViewResp, error) {
	v := video.NewVideoServiceClient(vc.conn)
	return v.View(context.Background(), req)
}

func (vc *videoClient) Judge(req *video.JudgeReq) (*video.JudgeResp, error) {
	v := video.NewVideoServiceClient(vc.conn)
	return v.Judge(context.Background(), req)
}

func (vc *videoClient) GetVideos(req *video.GetVideosReq) (*video.GetVideosResp, error) {
	v := video.NewVideoServiceClient(vc.conn)

	if req.Year < 1978 || req.Year > int64(time.Now().Year()) {
		return nil, erp.BadRequest
	}

	return v.GetVideos(context.Background(), req)
}
