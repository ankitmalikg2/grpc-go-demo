package main

import (
	"fmt"
	"log"
	"net"

	"./pb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("ankit malik")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}
	s := pb.Server{}
	grpcServer := grpc.NewServer()

	pb.RegisterChatServiceServer(grpcServer, &s)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

}
