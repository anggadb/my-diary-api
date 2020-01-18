package models

import (
	"MyDiaryApi/v1/database"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `gorm:"unique;not null" json:"email" form:"email"`
	Password string `gorm:"not null" json:"password" form:"password"`
	Username string `gorm:"not null" json:"username" form:"username"`
	Phone    string `gorm:"unique" json:"phone" form:"phone"`
	Address  string `gorm:"type:text" json:"address" form:"address"`
	Active   bool   `gorm:"default:true;not null" json:"active" form:"active"`
}

type FailLogin struct {
	message string
}

func (b *User) TableName() string {
	return "user"
}
func (b *User) BeforeCreate(scope *gorm.Scope) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(b.Password), 14)
	if err != nil {
		return err
	}
	scope.SetColumn("password", pass)
	return nil
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
func FindUserById(user *User, id string) (err error) {
	if err = database.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}
func FindUserByEmail(user *User, email string) (err error) {
	if err = database.DB.Where("email = ?", email).First(user).Error; err != nil {
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
	if err = database.DB.Where("id = ?", id).Unscoped().Delete(user).Error; err != nil {
		return err
	}
	return nil
}
func LoginUser(user *User, password string) (err error) {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}
