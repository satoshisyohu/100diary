package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type Config struct {
	Port   string
	DBHost string
	DBUser string
	DBPass string
}

func NewConfig() *Config {
	return &Config{
		Port:   os.Getenv("PORT"),
		DBHost: os.Getenv("DB_HOST"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
	}
}

func InitDb() (*gorm.DB, error) {
	cfg := NewConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/diary?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.Port)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
