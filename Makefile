gen:
	protoc --proto_path=proto proto/*.proto --go-grpc_out=require_unimplemented_servers=false:. --go_out=.

clean:
	rm pb/*.go

run:
	go run main.go

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

nginx:
	nginx -c nginx.conf -p $(CURDIR)

server1:
	go run cmd/server/main.go -port 50051

server2:
	go run cmd/server/main.go -port 50052

client-tls:
	go run cmd/client/main.go -address 0.0.0.0:8080 -tls

server1-tls:
	go run cmd/server/main.go -port 50051 -tls

server2-tls:
	go run cmd/server/main.go -port 50052 -tls

test:
	go test -cover -race ./...

cert:
	cd cert; ./gen.sh; cd ..

.PHONY: gen clean run server client test cert
