api/booksearch_grpc.pb.go api/booksearch.pb.go: api/booksearch.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/booksearch.proto

filldb:
	docker exec -i mysql mysql -uroot -p"pass" -D books < assets/sample_db.sql

tests:
	go test -v github.com/rprtr258/kvadotest/test

protoc: api/booksearch_grpc.pb.go api/booksearch.pb.go

dockerrun:
	docker-compose -f deployments/docker-compose.yml up --build

server:
	go run cmd/server/main.go

authorclient:
	go run cmd/client/main.go -author "Rowling"

titleclient:
	go run cmd/client/main.go -title "Python"

needleclient:
	go run cmd/client/main.go -needle "о"

authordockerclient:
	docker build -t client -f deployments/Dockerfile.client .
	docker run --network host client -author "Rowling"

titledockerclient:
	docker build -t client -f deployments/Dockerfile.client .
	docker run --network host client -title "Python"

needledockerclient:
	docker build -t client -f deployments/Dockerfile.client .
	docker run --network host client -needle "о"