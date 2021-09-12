package main

import (
	"context"
	"log"

	"github.com/wallacemachado/grpc-hello-go/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close() // após concluir o processo ele fecha a conexão

	client := pb.NewHelloServiceClient(conn)

	request := pb.HelloRequest{
		Name: "wallace",
	}

	res, err := client.Hello(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error during the execution: %v", err)
	}

	log.Println(res.Msg)

}
