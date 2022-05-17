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

	"github.com/rprtr258/kvadotest/internal"
	"github.com/rprtr258/kvadotest/internal/repositories"
	protobuf "github.com/rprtr258/kvadotest/pkg/api"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "Port to run server on")
	db   = flag.String("db", "root:pass@/books?multiStatements=true", "Database DSN to connect to")
)

func main() {
	flag.Parse()

	// Init tcp server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Init mysql repository
	booksRepo, err := repositories.NewMysqlRepository(*db)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer booksRepo.Close()
	// Init app server
	srv := internal.NewServer(booksRepo)
	// Init grpc server
	grpcServer := grpc.NewServer()
	protobuf.RegisterBookSearchServer(grpcServer, &srv)

	// Serve protobuf server
	go func() {
		log.Printf("server listening at %v", listener.Addr())
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Graceful stop on SIGINT
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	<-sig
	grpcServer.GracefulStop()
}
