package models

import (
	"chris_project/utils"
	"errors"
	"html"
	"strings"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
	Products []Product
}

//insert user into database
func (u *User) SaveUser() (*User, error) {
	err := DB.Create(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

//gorm hook - hash password
func (u *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := utils.Encrypt(u.Password)
	if err != nil {
		return nil
	}
	u.Password = string(hashedPassword)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	return nil
}

func (u *User) PrepareGive() {
	u.Password = ""
}

func GetUserByID(uid uint) (User, error) {
	var u User

	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found")
	}

	u.PrepareGive()
	return u, nil
}

func RegisterCheck(username, password string) (err error) {
	u := &User{}
	err = DB.Model(&User{}).Where("username = ?", username).First(&u).Error
	if err == nil {
		return errors.New("username already exist")
	}

	u.Username = username
	u.Password = password

	_, err = u.SaveUser()
	if err != nil {
		return errors.New("user insert error")
	}

	return
}

func LoginCheck(username, password string) (id uint, err error) {
	u := User{}

	err = DB.Model(&User{}).Where("username = ?", username).Take(&u).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return 0, errors.New("you haven't registered")
	}

	err = utils.Compare(u.Password, password)
	if err != nil {
		return 0, err
	}
	return u.ID, nil
}
