package database

import "gorm.io/gorm"

type FoodInventory struct {
	gorm.Model
	OffCode   string
	Favourite bool
	UserID    uint
	FoodDiary []FoodDiary
}

func FindOneFoodInventory(id string) (FoodInventory, error) {
	var f FoodInventory
	err := DB.Find(&f, id).Error
	return f, err
}

func DeleteFoodInventory(id string) error {
	var f FoodInventory
	if err := DB.First(&f, "id = ?", id).Error; err != nil {
		return err
	}
	return DB.Unscoped().Delete(f).Error
}

func FindManyFoodInventory(UserID uint) ([]FoodInventory, error) {
	var f []FoodInventory
	err := DB.Find(&f, "user_id = ?", UserID).Error
	return f, err
}
