package main

import (
	"MyDiaryApi/v1/database"
	"MyDiaryApi/v1/models"
	"MyDiaryApi/v1/routes"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
)

var err error

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Dokumen .env tidak ditemukan")
	}
}
func main() {
	database.DB, err = gorm.Open("postgres", database.DBURL(database.BuildDBConfig()))
	if err != nil {
		fmt.Println("DB Status : ", err)
	}
	defer database.DB.Close()
	database.DB.AutoMigrate(&models.User{})
	r := routes.UserRouter()
	r.Run(":8080")
}
