proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/*.proto

wire:
	cd pkg/di && wire

run:
	go run cmd/main.go