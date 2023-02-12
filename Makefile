.PHONY: server client clean api

server:
	go build -o bin/edged main.go

client:
	go build -o bin/edgectl client/main.go

clean:
	rm -rf bin

api:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/server.proto
