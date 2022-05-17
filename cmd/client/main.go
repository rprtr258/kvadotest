// Package main implements a client for BookSearch service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	protobuf "github.com/rprtr258/kvadotest/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	address = flag.String("addr", "localhost:50051", "Server address to connect to")
	author  = flag.String("author", "", "Author to search")
	title   = flag.String("title", "", "Title to search")
	needle  = flag.String("needle", "", "Books' content search needle")
)

// bool2int converts true to 1, false to 0
func bool2int(x bool) int {
	if x {
		return 1
	} else {
		return 0
	}
}

// validateFlags checks that exactly one search type is provided
func validateFlags() {
	notEmptyFlagsCount := bool2int(len(*author) != 0) + bool2int(len(*needle) != 0) + bool2int(len(*title) != 0)
	// Check that at least one search type is provided
	if notEmptyFlagsCount == 0 {
		log.Fatal("Provide one of flags: -author, -title or -needle")
	}
	// Check that no more than one search type is provided
	if notEmptyFlagsCount > 1 {
		log.Fatal("Too many flags provided")
	}
}

// doRequest sends protobuf request and get books list from response
func doRequest(
	ctx context.Context,
	client protobuf.BookSearchClient,
	searchRequest *protobuf.SearchRequest,
) ([]*protobuf.Book, error) {
	response, err := client.Search(ctx, searchRequest)
	if err != nil {
		return nil, err
	}
	return response.GetBooks(), nil
}

// printBooks prints books list
func printBooks(books []*protobuf.Book) {
	// If no books were found
	if len(books) == 0 {
		fmt.Println("No books found")
		return
	}
	fmt.Printf("Found %d books (showing only first 20 chars of content):\n", len(books))
	for _, book := range books {
		// Print book info
		fmt.Printf("Title: %s\n", book.Title)
		fmt.Println("Authors:")
		for _, author := range book.Authors {
			fmt.Printf("    %s\n", author)
		}
		// Get substring of first 20 runes to show
		contentShowedContentPart := string([]rune(book.Content)[:20])
		fmt.Printf("Content:\n%s\n\n", contentShowedContentPart)
	}
}

// run sends request to server and print result using protobuf client, timeouting after 1s
func run(client protobuf.BookSearchClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var (
		books []*protobuf.Book
		err   error
	)
	if len(*author) != 0 {
		// Send search by author name request
		books, err = doRequest(ctx, client, &protobuf.SearchRequest{Request: &protobuf.SearchRequest_ByAuthor{ByAuthor: *author}})
	} else if len(*needle) != 0 {
		// Send search by book content request
		books, err = doRequest(ctx, client, &protobuf.SearchRequest{Request: &protobuf.SearchRequest_ByContent{ByContent: *needle}})
	} else if len(*title) != 0 {
		// Send search by book title request
		books, err = doRequest(ctx, client, &protobuf.SearchRequest{Request: &protobuf.SearchRequest_ByTitle{ByTitle: *title}})
	}
	if err != nil {
		log.Fatalln(err)
	}
	printBooks(books)
}

func main() {
	flag.Parse()
	validateFlags()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create protobuf client
	client := protobuf.NewBookSearchClient(conn)

	run(client)
}
