package main

import (
	"MyDiaryApi/v1/database"
	"MyDiaryApi/v1/env"
	"MyDiaryApi/v1/models"
	"MyDiaryApi/v1/routes"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var err error

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
	routes.TestRouter(router)
	router.Run(env.Environment().Port)
}
