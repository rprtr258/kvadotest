// Package main implements a server for BookSearch service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"

	pb "github.com/rprtr258/kvadotest/api"
	"github.com/rprtr258/kvadotest/internal/repositories"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedBookSearchServer
	booksRepo repositories.BooksRepository
}

func (s *server) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchReply, error) {
	var (
		res []repositories.Book
		err error
	)
	switch req := in.Request.(type) {
	case *pb.SearchRequest_ByAuthor:
		log.Printf("Received by author request: %v", req.ByAuthor)
		res, err = s.booksRepo.SearchByAuthor(ctx, req.ByAuthor)
	case *pb.SearchRequest_ByContent:
		log.Printf("Received by content request %v", req.ByContent)
		res, err = s.booksRepo.SearchByContent(ctx, req.ByContent)
	case *pb.SearchRequest_ByTitle:
		log.Printf("Received by title request %v", req.ByTitle)
		res, err = s.booksRepo.SearchByTitle(ctx, req.ByTitle)
	}
	if err != nil {
		return nil, err
	}
	pbBooks := make([]*pb.Book, len(res))
	for i, book := range res {
		pbBooks[i] = &pb.Book{
			Authors: book.Authors,
			Title:   book.Title,
			Content: book.Content,
		}
	}
	return &pb.SearchReply{Books: pbBooks}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	booksRepo, err := repositories.NewMysqlRepository("root:pass@/books")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer booksRepo.Close()
	srv := server{
		booksRepo: booksRepo,
	}
	pb.RegisterBookSearchServer(s, &srv)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// TODO: graceful exit
}
