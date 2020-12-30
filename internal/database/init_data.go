package database

import (
	"os"
	"time"
)

type MealType struct {
	ID          string
	Name        string
	FoodDiaries []FoodDiary
}

type QuantityUnit struct {
	ID          uint
	Name        string
	FoodDiaries []FoodDiary
}

func InitData() error {
	// Admin/First user
	password, err := PasswordHash("123456")
	if err != nil {
		return err
	}

	user := User{
		Username: "admin",
		IsAdmin:  true,
		Password: password,
		UserSetting: UserSetting{
			FullName:  "Admin",
			Language:  "en",
			UseMetric: true,
			Birthday:  time.Date(1989, 01, 01, 22, 51, 48, 324359102, time.UTC),
		},
	}

	err = DB.FirstOrCreate(&user, 1).Error
	if err != nil {
		return err
	}

	// Food related
	m := []MealType{
		{ID: "B", Name: "Breakfast"},
		{ID: "L", Name: "Lunch"},
		{ID: "D", Name: "Dinner"},
		{ID: "S", Name: "Snacks"},
	}

	q := []QuantityUnit{
		{ID: 1, Name: "100 gr"},
		{ID: 2, Name: "1 g"},
		{ID: 3, Name: "3.5 ounce"},
		{ID: 4, Name: "1 ounce"},
	}

	for _, mt := range m {
		err := DB.Where(MealType{ID: mt.ID}).FirstOrCreate(&mt).Error
		if err != nil {
			return err
		}
	}

	for _, qu := range q {
		err := DB.Where(QuantityUnit{ID: qu.ID}).FirstOrCreate(&qu).Error
		if err != nil {
			return err
		}
	}

	// Create the directory for uploads
	path := "./data/img"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0700)
	}

	return nil
}
