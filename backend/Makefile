build:
	go build -o micro-reddit ./cmd/main.go
migrate:
	go run ./cmd/buildDb/main.go
seed:
	go run ./cmd/buildDb/populateDB/main.go
run: 
	go run ./cmd/main.go
all: migrate seed run
