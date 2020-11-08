package database

import (
	"time"

	"gorm.io/gorm"
)

type FoodDiary struct {
	gorm.Model
	MealTypeID      string  // Breakfast, Lunch, Dinner, Snacks, Water
	Quantity        float32 `gorm:"type:FLOAT"`
	Date            time.Time
	QuantityUnitID  uint
	Comments        string
	UserID          uint
	FoodInventoryID uint
}

func FindOneFoodDiary(id string) (FoodDiary, error) {
	var f FoodDiary
	err := DB.Find(&f, id).Error
	return f, err
}

func DeleteFoodDiary(id string) error {
	var f FoodDiary
	if err := DB.First(&f, "id = ?", id).Error; err != nil {
		return err
	}
	return DB.Unscoped().Delete(f).Error
}

func FindManyFoodDiary(UserID uint) ([]FoodDiary, error) {
	var f []FoodDiary
	err := DB.Find(&f, "user_id = ?", UserID).Error
	return f, err
}
