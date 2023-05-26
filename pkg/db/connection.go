package db

import (
	"fmt"

	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/config"
	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase(cfg *config.Config) (db *gorm.DB, err error) {

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database connection error:%s", err.Error())
	}

	err = db.AutoMigrate(domain.User{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate tables error:%s", err.Error())
	}
	return
}
