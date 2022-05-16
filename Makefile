api/booksearch_grpc.pb.go api/booksearch.pb.go: api/booksearch.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/booksearch.proto

protoc: api/booksearch_grpc.pb.go api/booksearch.pb.go

server:
	go run cmd/server/main.go

authorclient:
	go run cmd/client/main.go -author "Rowling"

titleclient:
	go run cmd/client/main.go -title "Python"

needleclient:
	go run cmd/client/main.go -needle "о"
