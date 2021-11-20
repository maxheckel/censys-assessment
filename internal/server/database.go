package server

import (
	"fmt"
	"github.com/maxheckel/censys-assessment/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func GetDB(config *config.Config) (*gorm.DB, error){
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort, config.DBSSLMode)
	count := 0
	var db *gorm.DB
	var err error
	// Very basic sleep / retry
	for count < 5{
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			return db, err
		}
		count++
		time.Sleep(5 * time.Second)
	}
	return db, err
}
