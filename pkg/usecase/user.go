package usecase

import (
	"context"

	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/domain"
	repo "github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/usecase/interfaces"
)

type userUseCase struct {
	repo repo.UserRepository
}

func NewUserUseCase(repo repo.UserRepository) interfaces.UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (c *userUseCase) SaveUser(ctx context.Context, user domain.User) (userId uint32, err error) {
	userId, err = c.repo.SaveUser(ctx, user)
	return
}
func (c *userUseCase) FindUserByEmail(ctx context.Context, email string) (user domain.User, err error) {
	user, err = c.repo.FindUserByEmail(ctx, email)
	return
}
