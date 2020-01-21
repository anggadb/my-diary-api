package models

import (
	"MyDiaryApi/v1/database"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Diary struct {
	gorm.Model
	UserID  uint
	Diary   string `gorm:"not null;type:text" json:"diary" form:"diary"`
	Private bool   `gorm:"not null;default:true" json:"private" form:"private"`
	Type    string `json:"type" form:"type"`
}

func (d *Diary) TableName() string {
	return "diary"
}
func (d *Diary) BeforeCreate(scope *gorm.Scope) (err error) {
	if d.Type == "draft" || d.Type == "published" {
		scope.SetColumn("type", d.Type)
		return nil
	} else {
		return err
	}
}
func CreateDiary(diary *Diary) (err error) {
	if err := database.DB.Create(diary).Error; err != nil {
		return err
	}
	return nil
}
func FindDiaryById(diary *Diary, id string) (err error) {
	if err := database.DB.Where("id = ?", id).First(diary).Error; err != nil {
		return err
	}
	return nil
}
func FindDiaryByUser(diary *Diary, user_id string) (err error) {
	if err := database.DB.Where("user_id = ? ", user_id).First(diary).Error; err != nil {
		return err
	}
	return nil
}
func FindAllDiaries(diary *[]Diary) (err error) {
	if err := database.DB.Find(diary).Error; err != nil {
		return err
	}
	return nil
}
func UpdateDiary(diary *Diary) (err error) {
	if err := database.DB.Save(diary).Error; err != nil {
		return nil
	}
	return nil
}
func DeleteDiary(diary *Diary) (err error) {
	if err := database.DB.Delete(diary).Error; err != nil {
		return err
	}
	return nil
}
