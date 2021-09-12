package main

import (
	"context"
	"log"
	"net"

	"github.com/wallacemachado/grpc-hello-go/pb"
	"google.golang.org/grpc"
)

type server struct {
}

// método previsto em proto/HelloService
func (*server) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	result := "Hello " + request.GetName()

	res := &pb.HelloResponse{
		Msg: result,
	}
	return res, nil
}

func main() {

	// subindo o servidor como se fosse o express do node
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to list %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterHelloServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to list %v", err)
	}
}
