package internal

import (
	"context"
	"log"

	pb "github.com/rprtr258/kvadotest/api"

	"github.com/rprtr258/kvadotest/internal/repositories"
)

type Server struct {
	pb.UnimplementedBookSearchServer
	booksRepo repositories.BooksRepository
}

func NewServer(booksRepo repositories.BooksRepository) Server {
	return Server{
		booksRepo: booksRepo,
	}
}

func (s *Server) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchReply, error) {
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
