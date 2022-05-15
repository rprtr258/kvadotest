// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/rprtr258/kvadotest/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr   = flag.String("addr", "localhost:50051", "the address to connect to")
	author = flag.String("author", "", "Author to search")
	needle = flag.String("needle", "", "Books' content search needle")
)

func doRequest(ctx context.Context, c pb.BookSearchClient, searchRequest *pb.SearchRequest) ([]*pb.Book, error) {
	r, err := c.Search(ctx, searchRequest)
	if err != nil {
		return nil, err
	}
	return r.GetBooks(), nil
}

func printBooks(books []*pb.Book) {
	if len(books) == 0 {
		fmt.Println("No books found")
		return
	}
	fmt.Printf("Found %d books (showing only first 10 chars of content):\n", len(books))
	for i, book := range books {
		fmt.Printf("%d. Author=%q Title=%q Content:\n%s", i, book.Author, book.Title, book.Content[:10])
	}
}

func main() {
	flag.Parse()
	if len(*author) == 0 && len(*needle) == 0 {
		log.Fatal("Neither -author nor -needle are provided")
	}
	if len(*author) != 0 && len(*needle) != 0 {
		log.Fatal("Both -author and -needle are provided, should be only one")
	}
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()
	c := pb.NewBookSearchClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if len(*author) != 0 {
		books, err := doRequest(ctx, c, &pb.SearchRequest{Request: &pb.SearchRequest_ByAuthor{ByAuthor: *author}})
		if err != nil {
			log.Fatalf("couldn't search by author: %v", err)
		}
		printBooks(books)
	} else if len(*needle) != 0 {
		books, err := doRequest(ctx, c, &pb.SearchRequest{Request: &pb.SearchRequest_ByAuthor{ByAuthor: *author}})
		if err != nil {
			log.Fatalf("couldn't search by content: %v", err)
		}
		printBooks(books)
	}
}
