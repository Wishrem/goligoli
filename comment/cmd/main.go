package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/wishrem/goligoli/comment/model"
	"github.com/wishrem/goligoli/comment/proto/pb"
	"github.com/wishrem/goligoli/comment/service"
	"github.com/wishrem/goligoli/logger"
	"github.com/wishrem/goligoli/pkg/conf"
	"github.com/wishrem/goligoli/pkg/util/snowflake"
	"google.golang.org/grpc"
)

var Log = logger.Log

func init() {
	model.Init()
	logger.Setup(logger.Debug, log.Default())
	snowflake.Init(1)
}

func main() {

	addr := fmt.Sprintf("localhost:%s", conf.App.CommentService.RpcPort)

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
	pb.RegisterCommentServiceServer(grpcServer, &service.CommentService{})
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
