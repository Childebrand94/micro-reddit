build:
	go build -o micro-reddit ./cmd/main.go && go build -o buildDb ./cmd/buildDb/main.go && go build -o populateDb ./cmd/buildDb/populateDb/main.go 
migrate:
	./buildDb
seed:
	./populateDb
run: 
	./micro-reddit	
all: migrate seed run
