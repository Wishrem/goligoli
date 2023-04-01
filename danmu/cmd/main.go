package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/wishrem/goligoli/pkg/conf"
	"github.com/wishrem/goligoli/pkg/util/snowflake"

	"github.com/wishrem/goligoli/danmu/model"
	"github.com/wishrem/goligoli/danmu/proto/pb"
	"github.com/wishrem/goligoli/danmu/service"
	"google.golang.org/grpc"
)

func init() {
	model.Init()
	snowflake.Init(3)
}

func main() {

	addr := fmt.Sprintf("localhost:%s", conf.App.DanmuService.RpcPort)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		etcdUnRegister(addr)
		if i, ok := s.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}
	}()

	err := etcdRegister(addr)
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(nil))
	pb.RegisterDanmuServiceServer(grpcServer, &service.DanmuService{})
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
