package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/domain"
)

type UserRepository interface {
	FindUserByEmail(ctx context.Context, email string) (user domain.User, err error)
	SaveUser(ctx context.Context, user domain.User) (userId uint32, err error)
}
