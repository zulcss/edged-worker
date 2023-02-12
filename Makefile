.PHONY: server client clean api all

all: clean api server client

server:
	go build -o bin/edged-worker main.go

client:
	go build -o bin/edgectl client/main.go

clean:
	rm -rf bin

api:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/server.proto
