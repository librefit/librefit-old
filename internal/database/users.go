package database

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string
	Password      string
	IsAdmin       bool
	UserSettingID uint
	UserSetting   UserSetting
	FoodDiaries   []FoodDiary
	FoodInventory []FoodInventory
	Measurement   []Measurement
	Upload        []Upload
}

func PasswordHash(password string) (string, error) {
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(passwordHash), nil
}

// FindOneUser search in the users table for a user.
func FindOneUser(condition interface{}) (User, error) {
	var u User
	err := DB.Preload("UserSetting").Where(condition).First(&u).Error
	return u, err
}

func (u *User) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
