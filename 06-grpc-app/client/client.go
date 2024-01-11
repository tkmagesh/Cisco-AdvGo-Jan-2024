package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"grpc-app/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
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
	// doServerStreaming(ctx, appServiceClient)
	// doServerStreamingWithCancellation(ctx, appServiceClient)
	doClientStreaming(ctx, appServiceClient)
}

func doClientStreaming(ctx context.Context, appServiceClient proto.AppServiceClient) {
	nos := []int64{3, 5, 4, 2, 6, 8, 7, 9, 1}
	clientStream, err := appServiceClient.CalculateAverage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, no := range nos {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Average req, no : ", no)
		req := &proto.AverageRequest{
			No: no,
		}
		if err := clientStream.Send(req); err != nil {
			log.Fatalln()
		}
	}
	if res, err := clientStream.CloseAndRecv(); err == nil {
		fmt.Println("average :", res.GetAverage())
	} else {
		log.Fatalln(err)
	}
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

func doServerStreamingWithCancellation(ctx context.Context, appServiceClient proto.AppServiceClient) {
	primeReq := &proto.PrimeRequest{
		Start: 2,
		End:   100,
	}

	// creating context with cancellation
	cancelCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	clientStream, err := appServiceClient.GeneratePrimes(cancelCtx, primeReq)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Hit ENTER to stop...!")
	go func() {
		fmt.Scanln()
		cancel()
	}()

LOOP:
	for {
		res, err := clientStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			if code := status.Code(err); code == codes.Canceled {
				fmt.Println("Cancellation initiated")
				break LOOP
			}
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
