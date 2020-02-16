package main

import (
	"MyDiaryApi/v1/database"
	"MyDiaryApi/v1/models"
	"MyDiaryApi/v1/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var err error

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Dokumen .env tidak ditemukan")
	}
}
func main() {
	database.DB, err = gorm.Open("postgres", database.DBURL(database.BuildDBConfig()))
	if err != nil {
		fmt.Println("DB Status : ", err)
	}
	defer database.DB.Close()
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Diary{})
	router := gin.Default()
	routes.DiaryRouter(router)
	routes.UserRouter(router)
	router.Run(os.Getenv("PORT"))
}
