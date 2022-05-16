// Package main implements a server for BookSearch service.
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/rprtr258/kvadotest/api"
	"github.com/rprtr258/kvadotest/internal"
	"github.com/rprtr258/kvadotest/internal/repositories"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "Server port")
	db   = flag.String("db", "root:pass@/books?multiStatements=true", "Database DSN")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	booksRepo, err := repositories.NewMysqlRepository(*db)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer booksRepo.Close()
	srv := internal.NewServer(booksRepo)
	pb.RegisterBookSearchServer(s, &srv)
	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	<-sig
	s.GracefulStop()
}
