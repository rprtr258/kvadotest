// App server implementation for use as grpc server
package internal

import (
	"context"
	"log"

	protobuf "github.com/rprtr258/kvadotest/pkg/api"

	"github.com/rprtr258/kvadotest/internal/repositories"
)

// App server to use as grpc server
type Server struct {
	protobuf.UnimplementedBookSearchServer
	booksRepo repositories.BooksRepository
}

// Create new app server using initialized books repository
func NewServer(booksRepo repositories.BooksRepository) Server {
	return Server{
		booksRepo: booksRepo,
	}
}

// Handle grpc search request
func (s *Server) Search(ctx context.Context, in *protobuf.SearchRequest) (*protobuf.SearchReply, error) {
	var (
		res []repositories.Book
		err error
	)

	// Log request and query books by request
	switch req := in.Request.(type) {
	case *protobuf.SearchRequest_ByAuthor:
		log.Printf("Received by author request: %v", req.ByAuthor)
		res, err = s.booksRepo.SearchByAuthor(ctx, req.ByAuthor)
	case *protobuf.SearchRequest_ByContent:
		log.Printf("Received by content request %v", req.ByContent)
		res, err = s.booksRepo.SearchByContent(ctx, req.ByContent)
	case *protobuf.SearchRequest_ByTitle:
		log.Printf("Received by title request %v", req.ByTitle)
		res, err = s.booksRepo.SearchByTitle(ctx, req.ByTitle)
	}
	if err != nil {
		return nil, err
	}

	// Make books array for grpc reply
	pbBooks := make([]*protobuf.Book, len(res))
	for i, book := range res {
		pbBooks[i] = &protobuf.Book{
			Authors: book.Authors,
			Title:   book.Title,
			Content: book.Content,
		}
	}
	return &protobuf.SearchReply{Books: pbBooks}, nil
}
