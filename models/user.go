package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID int `gorm:"primarykey`
	Name string `gorm:"required"`
	Email string `gorm:"required"`
}

func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}

	return nil
}

func GetUsers(db *gorm.DB, User *[]User) (err error) {
	err = db.Find(User).Error
	if err != nil {
		return err
	}

	return nil
}

func GetUser(db *gorm.DB, User *User, id string) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(db *gorm.DB, User *User) (err error) {
	db.Save(User)
	return nil
}

func DeleteUser(db *gorm.DB, User *User, id string) (err error) {
	db.Where("id = ?", id).Delete(User)
	return nil
}