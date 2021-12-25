package rpc

import (
	"fmt"
	"google.golang.org/grpc"
	"rating/idl/gen/movie"
)

var (
	ClientIP        = "127.0.0.1"
	MovieClientPort = 50001
)

var MovieClient movie.MovieServiceClient

func InitRpcClients() error {
	movieConn, err := grpc.Dial(fmt.Sprintf("%s:%d", ClientIP, MovieClientPort),
		grpc.WithInsecure())
	if err != nil {
		return err
	}
	MovieClient = movie.NewMovieServiceClient(movieConn)

	return nil
}
