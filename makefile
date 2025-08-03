MAIN = "cmd/main.go"

generate:
	protoc --go_out=. --go-grpc_out=. api/api.proto

make run:
	go run $(MAIN)
