// Package main implements a server for BookSearch service.
package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	_ "github.com/go-sql-driver/mysql"

	pb "github.com/rprtr258/kvadotest/api"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type Book struct {
	Title   string
	Content string
	Authors []string
}

type BooksRepository interface {
	SearchByAuthor(context.Context, string) ([]Book, error)
	SearchByTitle(context.Context, string) ([]Book, error)
	SearchByContent(context.Context, string) ([]Book, error)
}

type MysqlBooksRepository struct {
	db *sql.DB
}

func readBookRows(rows *sql.Rows) ([]Book, error) {
	res := make([]Book, 0)
	var (
		authors     []byte
		bookTitle   string
		bookContent string
		authorsList []string
	)
	for rows.Next() {
		rows.Scan(&authors, &bookTitle, &bookContent)
		if err := json.Unmarshal(authors, &authorsList); err != nil {
			return nil, err
		}
		res = append(res, Book{
			Authors: authorsList,
			Title:   bookTitle,
			Content: bookContent,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

func (repo MysqlBooksRepository) SearchByAuthor(ctx context.Context, req string) ([]Book, error) {
	var (
		rows *sql.Rows
		err  error
	)
	rows, err = repo.db.QueryContext(ctx, `
		WITH authors AS (
			SELECT book_id
			FROM book_list
			WHERE author_name LIKE CONCAT('%', ?, '%')
		)
		SELECT JSON_ARRAYAGG(author_name) AS authors, book_title, book_content
		FROM book_list
		GROUP BY book_id
		HAVING book_id IN (select * from authors);
	`, req)
	if err != nil {
		return nil, err
	}
	return readBookRows(rows)
}
func (repo MysqlBooksRepository) SearchByTitle(ctx context.Context, req string) ([]Book, error) {
	var (
		rows *sql.Rows
		err  error
	)
	rows, err = repo.db.QueryContext(ctx, `
			SELECT JSON_ARRAYAGG(author_name) AS authors, book_title, book_content
			FROM book_list
			WHERE book_title LIKE CONCAT('%', ?, '%')
			GROUP BY book_id;
		`, req)
	if err != nil {
		return nil, err
	}
	return readBookRows(rows)
}
func (repo MysqlBooksRepository) SearchByContent(ctx context.Context, req string) ([]Book, error) {
	var (
		rows *sql.Rows
		err  error
	)
	rows, err = repo.db.QueryContext(ctx, `
			SELECT JSON_ARRAYAGG(author_name) AS authors, book_title, book_content
			FROM book_list
			WHERE book_content LIKE CONCAT('%', ?, '%')
			GROUP BY book_id;
		`, req)
	if err != nil {
		return nil, err
	}
	return readBookRows(rows)
}

type server struct {
	pb.UnimplementedBookSearchServer
	booksRepo BooksRepository
}

func (s *server) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchReply, error) {
	var (
		res []Book
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
	db, err := sql.Open("mysql", "root:pass@/books")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	ctx := context.Background()
	if err = db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
	srv := server{booksRepo: MysqlBooksRepository{db}}
	pb.RegisterBookSearchServer(s, &srv)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// TODO: graceful exit
}
