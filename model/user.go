package model

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PhoneNum int64
	StuNum   int64
	RealName string
	QQNum    int64
	Username string
	Password string
	Portrait File `gorm:"foreignKey:RelationID"`
}

func (u *User) Create() error {
	if _, err := u.CheckByUP(); err == nil {
		err = errors.New("user exist")
		return err
	}
	if err := DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

//通过username和password查找
func (u *User) CheckByUP() (User, error) {
	var newu User
	err := DB.Where(&User{Username: u.Username, Password: u.Password}).First(&newu).Error
	if err != nil {
		return User{}, err
	}
	return newu, nil
}

//
func (u *User) GetUser() {
	DB.Model(&u).First(u)
	return
}

//
func (u *User) GetPortrait() {
	fs := GetFile(u.ID, 0)
	f := fs[0]
	if f.ID == 0 {
		f.Uri = "assets/user/default.png"
	}
	u.Portrait = f
}

//添加信息
func (u *User) Update() error {
	return DB.Save(u).Error
}
