package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"rating/config"
	"rating/idl/gen/rating"
	"rating/model"
)

func main() {
	if err := config.InitConfig(config.CfgFileMain); err != nil {
		panic(err)
	}
	if err := model.InitModel(); err != nil {
		panic(err)
	}

	var port int
	var ip string
	flag.IntVar(&port, "port", 50003, "port for the service")
	flag.StringVar(&ip, "ip", "127.0.0.1", "ip for the service")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	rating.RegisterRatingServiceServer(grpcServer, &Handler{})
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
