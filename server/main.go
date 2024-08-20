package main

import (
	"context"
	pb "grpc-buf-example/gen/proto"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAnswerServer
}

func (s *server) SayAddress(ctx context.Context, req *pb.PostRequest) (*pb.PostResponse, error) {

	// 郵便番号がint64型なので、stringに変換
	postcodestr := strconv.FormatInt(req.Requestpostcode, 10)

	return &pb.PostResponse{Responseaddress: "東京都新宿区四谷三栄町は、" + postcodestr}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAnswerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
