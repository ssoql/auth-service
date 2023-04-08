package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ssoql/auth-service/internal/app/entities"
)

type ClientDB struct {
	*gorm.DB
	options DatabaseOptions
}

type dbFactory struct {
	connStr string
}

type DatabaseOptions struct {
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabaseSchema   string
	DatabasePort     string
	DatabaseHost     string
}

func (r *dbFactory) GetClient() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(r.connStr), &gorm.Config{})
}

func InitializeDB(cfg DatabaseOptions) (*ClientDB, error) {
	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseSchema,
	)
	dbClient, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if err := dbClient.AutoMigrate(&entities.User{}); err != nil {
		return nil, err
	}

	return &ClientDB{dbClient, cfg}, nil
}
