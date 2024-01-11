package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"grpc-app/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial("localhost:50051", options)

	if err != nil {
		log.Fatalln(err)
	}

	appServiceClient := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()

	// doRequestResponse(ctx, appServiceClient)
	doServerStreaming(ctx, appServiceClient)
}

func doServerStreaming(ctx context.Context, appServiceClient proto.AppServiceClient) {
	primeReq := &proto.PrimeRequest{
		Start: 2,
		End:   100,
	}
	clientStream, err := appServiceClient.GeneratePrimes(ctx, primeReq)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := clientStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Prime No : %d\n", res.GetPrimeNo())
	}
}

func doRequestResponse(ctx context.Context, appServiceClient proto.AppServiceClient) {
	req := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	addResponse, err := appServiceClient.Add(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Add Result :", addResponse.GetResult())
}
