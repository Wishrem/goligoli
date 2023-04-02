package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/wishrem/goligoli/logger"
	"github.com/wishrem/goligoli/pkg/conf"
	"github.com/wishrem/goligoli/pkg/etcd"
	"github.com/wishrem/goligoli/pkg/util/snowflake"
	"github.com/wishrem/goligoli/video/model"
	"github.com/wishrem/goligoli/video/proto/pb"
	"github.com/wishrem/goligoli/video/service"
	"google.golang.org/grpc"
)

var Log = logger.Log

func init() {
	model.Init()
	logger.Setup(logger.Debug, log.Default())
	if err := os.Mkdir(conf.App.Res.VideoDir, 0666); err != nil && !os.IsExist(err) {
		panic(err)
	}
	snowflake.Init(1)
}

func main() {
	name := conf.App.VideoService.Name
	addr := fmt.Sprintf("localhost:%s", conf.App.VideoService.RpcPort)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		etcd.EtcdUnRegister(addr, name)
		if i, ok := s.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}
	}()

	err := etcd.EtcdRegister(addr, name)
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(nil))
	pb.RegisterVideoServiceServer(grpcServer, &service.VideoService{})
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
