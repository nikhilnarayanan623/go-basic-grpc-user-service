//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/api"
	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/api/service"
	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/config"
	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/db"
	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/repository"
	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/usecase"
)

func InitializeService(cfg *config.Config) (*api.Server, error) {

	wire.Build(
		db.InitializeDatabase,
		repository.NewUserRepository,
		usecase.NewUserUseCase,
		service.NewServiceServer,
		api.NewUserServer,
	)

	return &api.Server{}, nil
}
