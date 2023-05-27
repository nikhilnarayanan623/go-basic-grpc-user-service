package usecase

import (
	"context"
	"fmt"

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

func (c *userUseCase) GetUserProfile(ctx context.Context, userId uint32) (domain.User, error) {
	user, err := c.repo.FindUserById(ctx, userId)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get user details error:%s", err.Error())
	}
	return user, nil
}

func (c *userUseCase) SaveUser(ctx context.Context, user domain.User) (userId uint32, err error) {
	userId, err = c.repo.SaveUser(ctx, user)
	return
}
func (c *userUseCase) FindUserByEmail(ctx context.Context, email string) (user domain.User, err error) {
	user, err = c.repo.FindUserByEmail(ctx, email)
	return
}
