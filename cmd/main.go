package main

import (
	"log"

	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/config"
	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/di"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config error:%s", err.Error())
	}

	server, err := di.InitializeService(cfg)
	if err != nil {
		log.Fatalf("failed to intialize service error:%s", err.Error())
	}

	err = server.Start()
	if err != nil {
		log.Fatalf("failed to start error:%s", err.Error())
	}
}
