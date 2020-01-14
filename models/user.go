package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
)

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
	pass, e := bcrypt.GenerateFromPassword([]byte(b.Password), 14)
	if e != nil {
		fmt.Print(e)
		os.Exit(2)
	}
	fmt.Print(pass)
	return
}