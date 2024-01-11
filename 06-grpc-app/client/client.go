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
	// doClientStreaming(ctx, appServiceClient)
	doBiDiStreaming(ctx, appServiceClient)
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

/*
func doBiDiStreaming(ctx context.Context, appServiceClient proto.AppServiceClient) {
	persons := []*proto.PersonName{
		{FirstName: "Magesh", LastName: "Kuppan"},
		{FirstName: "Suresh", LastName: "Kannan"},
		{FirstName: "Ramesh", LastName: "Jayaraman"},
		{FirstName: "Rajesh", LastName: "Pandit"},
		{FirstName: "Ganesh", LastName: "Kumar"},
	}

	clientStream, err := appServiceClient.Greet(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	done := make(chan struct{})
	go func() {
		for {
			res, err := clientStream.Recv()
			if err != nil {
				close(done)
				log.Println(err)
				break
			}
			log.Println(res.GetMessage())
		}

	}()

	for _, person := range persons {
		req := &proto.GreetRequest{
			Person: person,
		}
		log.Printf("Sending Person : %s %s\n", person.FirstName, person.LastName)
		if err := clientStream.Send(req); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
	clientStream.CloseSend()

	<-done
}
*/

func doBiDiStreaming(ctx context.Context, appServiceClient proto.AppServiceClient) <-chan struct{} {
	clientStream, err := appServiceClient.Greet(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	go sendRequests(ctx, clientStream)
	done := make(chan struct{})
	go func() {
		fmt.Println("Press ENTER to cancel")
		fmt.Scanln()
		clientStream.CloseSend()
		close(done)
	}()
	go recvResponse(ctx, clientStream)
	return done
	// <-done
}

func sendRequests(ctx context.Context, clientStream proto.AppService_GreetClient) {
	persons := []*proto.PersonName{
		{FirstName: "Magesh", LastName: "Kuppan"},
		{FirstName: "Suresh", LastName: "Kannan"},
		{FirstName: "Ramesh", LastName: "Jayaraman"},
		{FirstName: "Rajesh", LastName: "Pandit"},
		{FirstName: "Ganesh", LastName: "Kumar"},
	}

	// done := make(chan struct{})

	for _, person := range persons {
		req := &proto.GreetRequest{
			Person: person,
		}
		log.Printf("Sending Person : %s %s\n", person.FirstName, person.LastName)
		if err := clientStream.Send(req); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func recvResponse(ctx context.Context, clientStream proto.AppService_GreetClient) {
	for {
		res, err := clientStream.Recv()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(res.GetMessage())
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
