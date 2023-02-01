package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/tkmagesh/IBM-AdvGo-Jan-2023/05-grpc-app/proto"
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
	client := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()

	// doRequestResponse(ctx, client)
	// doServerStreaming(ctx, client)
	// doClientStreaming(ctx, client)
	doBiDirectionalStreaming(ctx, client)
}

func doRequestResponse(ctx context.Context, client proto.AppServiceClient) {
	addRequest := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	addResponse, err := client.Add(timeoutCtx, addRequest)
	if err != nil {
		if code := status.Code(err); code == codes.DeadlineExceeded {
			fmt.Println("Timeout occured")
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("Add Response :", addResponse.GetResult())
}

func doServerStreaming(ctx context.Context, client proto.AppServiceClient) {
	var start int32 = 3
	var end int32 = 100
	req := &proto.PrimeRequest{
		Start: start,
		End:   end,
	}
	clientStream, err := client.FindPrimes(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := clientStream.Recv()
		if err == io.EOF {
			fmt.Println("All prime numbers received")
			break
		}
		if err != nil {
			grpcCode := status.Code(err)
			if grpcCode == codes.Unavailable {
				log.Fatalln("Server not available")
				break
			}
			log.Println(err)
			continue
		}
		fmt.Printf("Prime No : %d\n", res.GetPrimeNo())
	}
}

func doClientStreaming(ctx context.Context, client proto.AppServiceClient) {
	nos := []int32{3, 1, 4, 2, 5, 7}

	clientStream, err := client.CalculateAverage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, no := range nos {
		req := &proto.AverageRequest{
			No: no,
		}
		fmt.Printf("Sending data for Average, no : %d\n", no)
		if err := clientStream.Send(req); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("All data sent....")
	if res, err := clientStream.CloseAndRecv(); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("Average received : %d\n", res.GetResult())
	}
}

func doBiDirectionalStreaming(ctx context.Context, client proto.AppServiceClient) {
	personNames := []proto.PersonName{
		proto.PersonName{FirstName: "Magesh", LastName: "Kuppan"},
		proto.PersonName{FirstName: "Suresh", LastName: "Kannan"},
		proto.PersonName{FirstName: "Rajesh", LastName: "Pandit"},
		proto.PersonName{FirstName: "Ganesh", LastName: "Kumar"},
		proto.PersonName{FirstName: "Ramesh", LastName: "Jayaraman"},
	}
	clientStream, err := client.Greet(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	done := make(chan struct{})
	go func() {
		for {
			resp, err := clientStream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(resp.GetGreetMessage())
		}
		done <- struct{}{}
	}()
	for _, personName := range personNames {
		time.Sleep(500 * time.Millisecond)
		req := &proto.GreetRequest{
			Person: &personName,
		}
		fmt.Println("Sending : ", personName.FirstName, personName.LastName)
		clientStream.Send(req)
	}
	clientStream.CloseSend()
	<-done
}
