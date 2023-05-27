package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/domain"
)

type UserUseCase interface {
	SaveUser(ctx context.Context, user domain.User) (userId uint32, err error)
	FindUserByEmail(ctx context.Context, email string) (domain.User, error)
	GetUserProfile(ctx context.Context, userId uint32) (domain.User, error)
}
