package main

import (
	"log"
	"net"

	"github.com/rschio/grpc-poc/auth"
	"github.com/rschio/grpc-poc/interceptor"
	"github.com/rschio/grpc-poc/server"
	"github.com/rschio/grpc-poc/tracker/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	l, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatal(err)
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			interceptor.UnaryAuth(auth.MyAuther{}),
			interceptor.UnaryLog,
		),
		grpc.ChainStreamInterceptor(
			interceptor.StreamAuth(auth.MyAuther{}),
			interceptor.StreamLog,
		),
	}

	grpcSrv := grpc.NewServer(opts...)
	proto.RegisterTrackerServer(grpcSrv, new(server.Server))

	// Used by Evans...
	reflection.Register(grpcSrv)

	err = grpcSrv.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
