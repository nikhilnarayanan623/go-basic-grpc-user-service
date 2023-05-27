proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/*.proto

wire:
	cd pkg/di && wire

run:
	go run cmd/main.go

docker-build:
	docker build -t nikhil382/go-basic-grpc-user-service .

docker-up:
	docker-compose up 

docker-down:
	docker-compose down