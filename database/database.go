package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
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
		User:     os.Getenv("DB_USER"),
		Name:     os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Sslmode:  os.Getenv("DB_SSLMODE"),
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
