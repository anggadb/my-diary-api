package models

import (
	"MyDiaryApi/v1/database"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"os"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Username string `gorm:"not null" json:"username"`
	Phone    string `gorm:"unique" json:"phone"`
	Address  string `gorm:"type:text" json:"address"`
}

func (b *User) TableName() string {
	return "user"
}
func (b *User) BeforeCreate() (err error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(b.Password), 14)
	if err != nil {
		fmt.Print(err)
		os.Exit(2)
	}
	fmt.Print(pass)
	return
}
func CreateUser(user *User) (err error) {
	if err = database.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
func FindAllUsers(user *[]User) (err error) {
	if err = database.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}
func FindUser(user *User, id string) (err error) {
	if err = database.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}
func UpdateUser(user *User, id string) (err error) {
	if err = database.DB.Where("id = ?", id).Save(user).Error; err != nil {
		return err
	}
	return nil
}
func DeleteUser(user *User, id string) (err error) {
	if err = database.DB.Where("id = ?", id).Delete(user).Error; err != nil {
		return err
	}
	return nil
}
