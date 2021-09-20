gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:.

clean:
	rm -rf pb/*.go

server:
	go run cmd/server/main.go -port 9999

client:
	go run cmd/client/main.go -address 0.0.0.0:9999

test:
	go test -cover -race ./...