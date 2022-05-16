// Package main implements a client for BookSearch service.
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
	title  = flag.String("title", "", "Title to search")
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
	fmt.Printf("Found %d books (showing only first 20 chars of content):\n", len(books))
	for _, book := range books {
		fmt.Printf("Title: %s\n", book.Title)
		fmt.Println("Authors:")
		for _, author := range book.Authors {
			fmt.Printf("    %s\n", author)
		}
		fmt.Printf("Content:\n%s\n\n", string([]rune(book.Content)[:20]))
	}
}

func bool2int(x bool) int {
	if x {
		return 1
	} else {
		return 0
	}
}

func main() {
	flag.Parse()
	notEmptyFlags := bool2int(len(*author) != 0) + bool2int(len(*needle) != 0) + bool2int(len(*title) != 0)
	if notEmptyFlags == 0 {
		log.Fatal("Provide one of flags: -author, -title or -needle")
	}
	if notEmptyFlags > 1 {
		log.Fatal("Too many flags provided")
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
	var books []*pb.Book
	if len(*author) != 0 {
		books, err = doRequest(ctx, c, &pb.SearchRequest{Request: &pb.SearchRequest_ByAuthor{ByAuthor: *author}})
		if err != nil {
			log.Fatalf("couldn't search by author: %v", err)
		}
	} else if len(*needle) != 0 {
		books, err = doRequest(ctx, c, &pb.SearchRequest{Request: &pb.SearchRequest_ByContent{ByContent: *needle}})
		if err != nil {
			log.Fatalf("couldn't search by content: %v", err)
		}
	} else if len(*title) != 0 {
		books, err = doRequest(ctx, c, &pb.SearchRequest{Request: &pb.SearchRequest_ByTitle{ByTitle: *title}})
		if err != nil {
			log.Fatalf("couldn't search by title: %v", err)
		}
	}
	printBooks(books)
}
