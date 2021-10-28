package database

import (
	"MyDiaryApi/v1/env"
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func DBConfig() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		env.Environment().DBHost,
		env.Environment().DBPort,
		env.Environment().DBUser,
		env.Environment().DBPassword,
		env.Environment().DBName,
		env.Environment().DBSSLMode,
	)
}
