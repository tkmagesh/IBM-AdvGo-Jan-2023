package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/tkmagesh/IBM-AdvGo-Jan-2023/05-grpc-app/proto"
	"google.golang.org/grpc"
)

type AppServiceImpl struct {
	proto.UnimplementedAppServiceServer
}

func (asi *AppServiceImpl) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	result := x + y
	log.Printf("Processing %d and %d. Returning %d\n", x, y, result)
	time.Sleep(10 * time.Second)
	res := &proto.AddResponse{
		Result: result,
	}
	return res, nil
}

func (asi *AppServiceImpl) FindPrimes(req *proto.PrimeRequest, serverStream proto.AppService_FindPrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	fmt.Printf("Processing FindPrime request with start = %d and end = %d\n", start, end)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			res := &proto.PrimeResponse{
				PrimeNo: no,
			}
			fmt.Printf("Sending : %d\n", no)
			if err := serverStream.Send(res); err != nil {
				log.Println(err)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
	fmt.Println("All prime numbers sent!")
	return nil
}

func isPrime(no int32) bool {
	for i := int32(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	//hosting the service
	asi := &AppServiceImpl{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, asi)
	grpcServer.Serve(listener)
}
