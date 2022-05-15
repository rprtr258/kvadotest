// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/rprtr258/kvadotest/api"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnsafeBookSearchServer
}

func (s *server) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchReply, error) {
	switch req := in.Request.(type) {
	case *pb.SearchRequest_ByAuthor:
		log.Printf("Received by author request: %v", req.ByAuthor)
	case *pb.SearchRequest_ByContent:
		log.Printf("Received: by content request %v", req.ByContent)
	}
	return &pb.SearchReply{Books: []*pb.Book{}}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBookSearchServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
