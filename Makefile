pkg/api/booksearch_grpc.pb.go pkg/api/booksearch.pb.go: api/booksearch.proto
	protoc --go_out=pkg --go_opt=paths=source_relative --go-grpc_out=pkg --go-grpc_opt=paths=source_relative api/booksearch.proto

# Generate grpc files from protobuf description
protoc: pkg/api/booksearch_grpc.pb.go pkg/api/booksearch.pb.go

# Generate mocks
mockgen:
	go generate ./...

# Run tests
tests: mockgen protoc
	go test -v -cover ./...

# Fill dockerized database with sample data
filldb:
	docker exec -i mysql mysql -uroot -p"pass" -D books < assets/sample_db.sql

# Run dockerized database and server
dockerrun: protoc
	docker-compose -f deployments/docker-compose.yml up --build

# Run dockerized database
dbdocker:
	docker-compose -f deployments/docker-compose.yml up db

# Run local server
server: protoc
	go run cmd/server/main.go

# Run client with sample author request
authorclient: protoc
	go run cmd/client/main.go -author "Rowling"

# Run client with sample title request
titleclient: protoc
	go run cmd/client/main.go -title "Python"

# Run client with sample content request
needleclient: protoc
	go run cmd/client/main.go -needle "о"

# Run dockerized client with sample author request
authordockerclient: protoc
	docker build -t client -f deployments/Dockerfile.client .
	docker run --network host client -author "Rowling"

# Run dockerized client with sample title request
titledockerclient: protoc
	docker build -t client -f deployments/Dockerfile.client .
	docker run --network host client -title "Python"

# Run dockerized client with sample content request
needledockerclient: protoc
	docker build -t client -f deployments/Dockerfile.client .
	docker run --network host client -needle "о"