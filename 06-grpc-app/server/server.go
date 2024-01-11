package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type AppServiceImpl struct {
	proto.UnimplementedAppServiceServer
}

func (as *AppServiceImpl) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	fmt.Printf("Processing Add Request with x = %d and y = %d\n", x, y)
	result := x + y
	fmt.Printf("Responding with result : %d\n", result)
	res := &proto.AddResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	asi := &AppServiceImpl{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, asi)
	grpcServer.Serve(listener)
}
