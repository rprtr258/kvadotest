pkg/helloworld_grpc.pb.go pkg/helloworld.pb.go: pkg/helloworld.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/helloworld.proto