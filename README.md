[![Go Report Card](https://goreportcard.com/badge/github.com/rprtr258/kvadotest)](https://goreportcard.com/report/github.com/rprtr258/kvadotest)

# Test task for kvado.ru

gRPC server to search books by author, title or content.

Sample dataset provided by [kaggle](https://www.kaggle.com/datasets/jealousleopard/goodreadsbooks)

## How to run

```bash
make dockerrun # Run database and server in docker
make filldb # Populate database with sample data
```
When everything is set up and correct you can run requests using dockerized or local client:

Using dockerized client:
```bash
make authordockerclient # sample search by author
make titledockerclient # sample search by title
make needledockerclient # sample search by content

# custom search
docker build -t client -f deployments/Dockerfile.client .
docker run --network host client <CUSTOM SEARCH>
```

Using local client
```bash
make authorclient # sample search by author
make titleclient # sample search by title
make needleclient # sample search by content

# custom search
go run cmd/client/main.go <CUSTOM SEARCH>
```

Where `<CUSTOM SEARCH>` is one of:
- `-author "AUTHOR"` to search by author
- `-title "TITLE"` to search by title
- `-needle "NEEDLE"` to search by content

You can find some other commands in [Makefile](Makefile)