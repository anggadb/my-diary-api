package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Name     string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Name:     "my_diary",
		Password: "Password",
	}
	return &dbConfig
}
func DBURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)
}
