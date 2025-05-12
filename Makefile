.PHONY: run
run:
	go build -o build/server cmd/server/main.go
	./build/server


.PHONY: generate
generate:
	rm -rf pkg/pb
	mkdir -p pkg/pb
	protoc --proto_path=api/ --go_out=pkg/pb --go-grpc_out=pkg/pb api/api.proto

.DEFAULT_GOAL := run
