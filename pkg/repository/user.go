package repository

import (
	"context"

	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/domain"
	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type userDatabase struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userDatabase{
		db: db,
	}
}

func (c *userDatabase) FindUserById(ctx context.Context, userId uint32) (user domain.User, err error) {
	query := `SELECT * FROM users WHERE id = $1`
	err = c.db.Raw(query, userId).Scan(&user).Error
	return
}

func (c *userDatabase) FindUserByEmail(ctx context.Context, email string) (user domain.User, err error) {
	query := `SELECT * FROM users WHERE email = $1`
	err = c.db.Raw(query, email).Scan(&user).Error
	return
}
func (c *userDatabase) SaveUser(ctx context.Context, user domain.User) (userId uint32, err error) {
	query := `INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4 ) RETURNING id AS user_id`
	err = c.db.Raw(query, user.FirstName, user.LastName, user.Email, user.Password).Scan(&userId).Error
	return
}
