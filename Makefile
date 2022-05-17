api/booksearch_grpc.pb.go api/booksearch.pb.go: api/booksearch.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/booksearch.proto

# Generate grpc files from protobuf description
protoc: api/booksearch_grpc.pb.go api/booksearch.pb.go

# Fill dockerized database with sample data
filldb:
	docker exec -i mysql mysql -uroot -p"pass" -D books < assets/sample_db.sql

# Run tests
tests:
	go test -v github.com/rprtr258/kvadotest/test

# Run dockerized database and server
dockerrun:
	docker-compose -f deployments/docker-compose.yml up --build

# Run local server
server:
	go run cmd/server/main.go

# Run client with sample author request
authorclient:
	go run cmd/client/main.go -author "Rowling"

# Run client with sample title request
titleclient:
	go run cmd/client/main.go -title "Python"

# Run client with sample content request
needleclient:
	go run cmd/client/main.go -needle "о"

# Run dockerized client with sample author request
authordockerclient:
	docker build -t client -f deployments/Dockerfile.client .
	docker run --network host client -author "Rowling"

# Run dockerized client with sample title request
titledockerclient:
	docker build -t client -f deployments/Dockerfile.client .
	docker run --network host client -title "Python"

# Run dockerized client with sample content request
needledockerclient:
	docker build -t client -f deployments/Dockerfile.client .
	docker run --network host client -needle "о"