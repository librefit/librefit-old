package database

import (
	"gorm.io/gorm"
)

type UserSetting struct {
	gorm.Model
	Country       string // Used for searching in OpenFoodFacts
	Timezone      string
	Language      string
	Theme         string
	FullName      string
	Email         string
	Birthday      string
	UseMetric     bool
	Sex           string
	Height        float32
	BodyFat       string
	ActivityLevel string
}

// FindOneUser search in the users table for a user.
func FindOneSetting(condition interface{}) (UserSetting, error) {
	var u UserSetting
	err := DB.Where(condition).First(&u).Error
	return u, err
}
