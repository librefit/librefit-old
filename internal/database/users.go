package database

import (
	"errors"

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
}

// CreateAdmin creates the initial user if doesn't exist.
func createAdmin() error {
	user := User{
		Username: "admin",
		IsAdmin:  true,
		UserSetting: UserSetting{
			Language:  "en",
			UseMetric: true,
		},
	}

	err := user.setPassword("123456")
	if err != nil {
		return err
	}

	err = DB.FirstOrCreate(&user, 1).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *User) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty!")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(passwordHash)
	return nil
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
