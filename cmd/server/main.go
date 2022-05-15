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

type server struct {
	pb.UnimplementedBookSearchServer
	db *sql.DB
}

func readBookRows(rows *sql.Rows) ([]*pb.Book, error) {
	res := make([]*pb.Book, 0)
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
		res = append(res, &pb.Book{
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

func (s *server) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchReply, error) {
	var res []*pb.Book
	switch req := in.Request.(type) {
	case *pb.SearchRequest_ByAuthor:
		log.Printf("Received by author request: %v", req.ByAuthor)
		rows, err := s.db.QueryContext(ctx, `
			WITH authors AS (
				SELECT book_id
				FROM book_list
				WHERE author_name LIKE CONCAT('%', ?, '%')
			)
			SELECT JSON_ARRAYAGG(author_name) AS authors, book_title, book_content
			FROM book_list
			GROUP BY book_id
			HAVING book_id IN (select * from authors);
		`, req.ByAuthor)
		if err != nil {
			return nil, err
		}
		if res, err = readBookRows(rows); err != nil {
			return nil, err
		}
	case *pb.SearchRequest_ByContent:
		log.Printf("Received by content request %v", req.ByContent)
		rows, err := s.db.QueryContext(ctx, `
			SELECT JSON_ARRAYAGG(author_name) AS authors, book_title, book_content
			FROM book_list
			GROUP BY book_id
			HAVING book_id IN (
				SELECT id
				FROM book
				WHERE content LIKE CONCAT('%', ?, '%')
			);
		`, req.ByContent)
		if err != nil {
			return nil, err
		}
		if res, err = readBookRows(rows); err != nil {
			return nil, err
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
	srv := server{db: db}
	pb.RegisterBookSearchServer(s, &srv)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// TODO: graceful exit
}
