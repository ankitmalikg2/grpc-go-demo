package main

import (
	"context"
	"log"

	"./pb"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("error occured with conncecting ", err)
	}
	defer conn.Close()

	msg := pb.Message{
		Body: "this is client message",
	}
	c := pb.NewChatServiceClient(conn)

	response, err := c.SayHello(context.Background(), &msg)
	if err != nil {
		log.Fatal("error in response: ", err)
	}
	log.Println(response)

}
