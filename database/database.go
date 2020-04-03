package database

import (
	"MyDiaryApi/v1/env"
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	User     string
	Name     string
	Password string
	Sslmode  string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		User:     env.Environment().DBUser,
		Name:     env.Environment().DBName,
		Password: env.Environment().DBPassword,
		Sslmode:  env.Environment().DBSSLMode,
	}
	return &dbConfig
}
func DBURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Sslmode,
	)
}
