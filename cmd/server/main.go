// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	pb "github.com/rprtr258/kvadotest/api"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
	data = []*pb.Book{
		{
			Title:   "Harry Potter and the Half-Blood Prince (Harry Potter  #6)",
			Authors: []string{"J.K. Rowling", "Mary GrandPré"},
			Content: "ABOBAABOBAABOBA",
		}, {
			Title:   "The Ultimate Hitchhiker's Guide: Five Complete Novels and One Story (Hitchhiker's Guide to the Galaxy  #1-5)",
			Authors: []string{"Douglas Adams"},
			Content: "zzzzzzzzzzzzzzzzzzzzzz",
		},
	}
)

type server struct {
	pb.UnsafeBookSearchServer
}

func (s *server) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchReply, error) {
	res := make([]*pb.Book, 0)
	switch req := in.Request.(type) {
	case *pb.SearchRequest_ByAuthor:
		log.Printf("Received by author request: %v", req.ByAuthor)
		for _, book := range data {
			for _, author := range book.Authors {
				if strings.Contains(author, req.ByAuthor) {
					res = append(res, book)
					break
				}
			}
		}
	case *pb.SearchRequest_ByContent:
		log.Printf("Received by content request %v", req.ByContent)
		for _, book := range data {
			if strings.Contains(book.Content, req.ByContent) {
				res = append(res, book)
			}
		}
	}
	return &pb.SearchReply{Books: res}, nil
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
