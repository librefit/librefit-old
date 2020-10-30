package database

import "gorm.io/gorm"

type FoodDiary struct {
	gorm.Model
	MealType        string // Breakfast, Lunch, Dinner, Snacks, Water
	Quantity        float32
	QuantityUnit    string
	Servings        float32
	Comments        string
	UserID          uint
	FoodInventoryID uint
}

type FoodInventory struct {
	gorm.Model
	OffCode   string
	Favourite bool
	UserID    uint
}
