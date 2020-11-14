package database

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string
	Password       string
	IsAdmin        bool
	Language       string
	Country        string // Used for searching in OpenFoodFacts
	Timezone       string
	Theme          string
	FullName       string
	Email          string
	Birthday       string
	PreferredUnits string
	Sex            string
	Height         float32
	BodyFat        string
	ActivityLevel  string

	Extra         string
	FoodDiaries   []FoodDiary
	FoodInventory []FoodInventory
	Measurement   []Measurement
	TestAle       datatypes.JSON
}

// CreateAdmin creates the initial user if doesn't exist.
func createAdmin() error {
	user := User{
		Username: "admin",
		IsAdmin:  true,
		Language: "en",
	}

	err := user.setPassword("123456")
	if err != nil {
		return err
	}

	err = DB.Where(User{Username: "admin"}).FirstOrCreate(&user).Error
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
	err := DB.Where(condition).First(&u).Error
	return u, err
}

func (u *User) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func (u *User) AddJSON() error {
	u.TestAle = datatypes.JSON([]byte(`{"name": "jinzhu", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orgaaaaa"}}`))

	DB.Save(&u)

	return nil
}
