package database

import "gorm.io/gorm"

type FoodDiary struct {
	gorm.Model
	Meal   string // Breakfast, Lunch, Dinner, Snacks, Water
	Notes  string
	UserID uint
}
