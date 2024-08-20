package main

import (
	"context"
	"log"
	"time"

	pb "grpc-buf-example/gen/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAnswerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayAddress(ctx, &pb.PostRequest{Requestpostcode: 1600008})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("住所と郵便番号：%s", r.GetResponseaddress())
}
