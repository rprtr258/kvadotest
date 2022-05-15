api/helloworld_grpc.pb.go api/helloworld.pb.go: api/helloworld.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/helloworld.proto

build: api/helloworld_grpc.pb.go api/helloworld.pb.go
