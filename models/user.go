package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"size:64,index"`
	Password string `gorm:"size:255"`
	IsAdmin  bool
}

func CheckUserAvailability(email string) bool {
	var user User
	user.ID = 0
	err := DB.Where("email = ?", email).First(&user)
	return err != nil
}

func UserCreate(email string, password string) *User {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}

	user := User{Email: email, Password: string(hashed), IsAdmin: false}
	return &user
}

func AdminCreate(email string, password string) *User {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}

	user := User{Email: email, Password: string(hashed), IsAdmin: true}
	return &user
}

func UserMatchPassword(email string, password string) *User {
	var user User
	DB.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return &user
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return &User{}
	} else {
		return &user
	}
}

func UserFromId(id uint) *User {
	var user User
	DB.Where("id = ?", id).First(&user)
	return &user
}
