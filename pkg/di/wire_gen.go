// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/api"
	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/api/service"
	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/config"
	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/db"
	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/repository"
	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/usecase"
)

// Injectors from wire.go:

func InitializeService(cfg *config.Config) (*api.Server, error) {
	gormDB, err := db.InitializeDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userServiceServer := service.NewServiceServer(userUseCase)
	server, err := api.NewUserServer(cfg, userServiceServer)
	if err != nil {
		return nil, err
	}
	return server, nil
}
