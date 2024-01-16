package main

import (
	"github.com/rs/zerolog"
	"go-grpc-service/api"
	"google.golang.org/grpc"
	"net"
	"os"
	"time"
)

func main() {
	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime},
	).Level(zerolog.TraceLevel).With().Timestamp().Caller().Logger()

	server, _ := net.Listen("tcp", "localhost:8080")
	grpcServer := grpc.NewServer()
	api.RegisterMovieServiceServer(grpcServer, &movieServiceServer{})
	logger.Info().Msg("gRPC server has been started.")
	panic(grpcServer.Serve(server))
}
