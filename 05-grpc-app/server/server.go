package main

import (
	"context"
	"fmt"
	"io"
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

func (asi *AppServiceImpl) CalculateAverage(serverStream proto.AppService_CalculateAverageServer) error {
	var sum int32 = 0
	var count int32 = 0
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			fmt.Println("Received all the requests")
			break
		}
		if err != nil {
			log.Fatalln()
		}
		sum += req.GetNo()
		count++
	}
	avg := sum / count
	res := &proto.AverageResponse{
		Result: avg,
	}
	if err := serverStream.SendAndClose(res); err != nil {
		log.Fatalln(err)
	}
	return nil
}

func (s *AppServiceImpl) Greet(serverStream proto.AppService_GreetServer) error {
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		person := req.GetPerson()
		msg := fmt.Sprintf("Hi %s %s!", person.GetFirstName(), person.GetLastName())
		resp := &proto.GreetResponse{
			GreetMessage: msg,
		}
		time.Sleep(500 * time.Millisecond)
		e := serverStream.Send(resp)
		if e != nil {
			log.Fatalln(err)
		}
	}
	return nil
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
