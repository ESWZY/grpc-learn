gen:
	protoc --proto_path=proto proto/*.proto --go-grpc_out=.

clean:
	rm pb/*.go

run:
	go run main.go

test:
	go test -cover -race ./...
