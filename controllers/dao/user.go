package dao

import (
	"errors"
	"go_jwt/db"

	"gorm.io/gorm"
)

var dbIns *gorm.DB

type User struct {
	Id       uint `gorm:"primary_key"`
	UserName string
	Password string
}

func Init(){
	dbIns = db.InitDB()
}

func (u *User) Add() (userId uint, err error) {
	if u.UserName == "" || u.Password == "" {
		return 0, errors.New("user_name or password empty!")
	}
	oUser := u.CheckHaveUserName(u.UserName)
	if oUser.Id > 0 {
		return oUser.Id, nil
	}
	if err = dbIns.Create(&u).Error; err != nil {
		return 0, err
	}

	return u.Id, nil
}

func (u *User) CheckHaveUserName(userName string) (data User) {
	dbIns.Where("username = ?", userName).Take(&data)
	return
}