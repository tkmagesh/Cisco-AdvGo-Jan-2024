package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"log"
	"net"
	"time"

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

func (as *AppServiceImpl) GeneratePrimes(req *proto.PrimeRequest, serverStream proto.AppService_GeneratePrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	fmt.Printf("GeneratePrime req received for start = %d and end = %d\n", start, end)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			fmt.Printf("Sending prime no : %d\n", no)
			res := &proto.PrimeResponse{
				PrimeNo: no,
			}
			serverStream.Send(res)
			time.Sleep(500 * time.Millisecond)
		}
	}
	return nil
}

func isPrime(no int64) bool {
	for i := int64(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
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
