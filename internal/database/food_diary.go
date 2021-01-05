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
	FoodInventory   FoodInventory
}

func FindOneFoodDiary(id string) (FoodDiary, error) {
	var f FoodDiary
	err := DB.Preload("FoodInventory.FoodInventoryImg.Upload").Preload("FoodInventory.FoodInventoryImg").Preload("FoodInventory").Find(&f, id).Error
	return f, err
}

func DeleteFoodDiary(id string) error {
	var f FoodDiary
	if err := DB.First(&f, "id = ?", id).Error; err != nil {
		return err
	}
	return DB.Unscoped().Delete(f).Error
}

func FindManyFoodDiary(UserID uint, start time.Time, end time.Time) ([]FoodDiary, error) {
	var f []FoodDiary
	// Custom Preloading to get soft-deleted items. When we delete product from inventory we still want
	// it logged in the diary
	err := DB.Preload("FoodInventory", func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}).Where("date BETWEEN ? AND ?", start, end).Find(&f, "user_id = ?", UserID).Error
	return f, err
}
