package main

import (
	"go-grpc-service/api"
	"google.golang.org/grpc"
	"net"
)

func main() {
	server, _ := net.Listen("tcp", "localhost:8080")
	grpcServer := grpc.NewServer()
	api.RegisterMovieServiceServer(grpcServer, &movieServiceServer{})
	panic(grpcServer.Serve(server))
}
