package database

import (
	"fmt"
	"log"

	"github.com/bananafried525/wallet/user/config"
	"github.com/bananafried525/wallet/user/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Connection struct {
	*gorm.DB
}

func New(cfg *config.Config) *Connection {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.UserName,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Database,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}

	db.Migrator().HasTable(&entity.User{})

	return &Connection{db}
}
