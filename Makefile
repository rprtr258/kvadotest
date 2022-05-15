api/booksearch_grpc.pb.go api/booksearch.pb.go: api/booksearch.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/booksearch.proto

build: api/booksearch_grpc.pb.go api/booksearch.pb.go

server:
	go run cmd/server/main.go

client:
	go run cmd/client/main.go
