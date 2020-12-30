package database

import (
	"time"

	"gorm.io/gorm"
)

type UserSetting struct {
	gorm.Model
	Country       string // Used for searching in OpenFoodFacts
	Timezone      string
	Avatar        string
	Language      string
	Theme         string
	FullName      string
	Email         string
	Birthday      time.Time
	UseMetric     bool
	Sex           string
	Height        float32
	BodyFat       string
	ActivityLevel string
	UserID        uint
}

// FindOneUser search in the users table for a user.
func FindOneSetting(UserID int) (UserSetting, error) {
	var u UserSetting
	err := DB.Find(&u, "user_id = ?", UserID).Error
	return u, err
}
