package database

import (
	openfoodfacts "github.com/mogaal/openfoodfacts-go"
	"gorm.io/gorm"
)

type FoodInventory struct {
	gorm.Model
	OffCode   string
	Favourite bool
	UserID    uint
	FoodDiary []FoodDiary
	// Important values to save so we're not always querying OFF
	Calories           float64
	FatTotal           float64
	FatSaturated       float64
	FatPolyunsaturated float64
	FatMonounsaturated float64
	FatTrans           float64
	FatCholesterol     float64
	Sodium             float64
	Potassium          float64
	Carbs              float64
	Fibers             float64
	Sugars             float64
	Proteins           float64
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

func FindOrCreateFoodInventory(u uint, o string) (*FoodInventory, error) {
	var r FoodInventory

	if err := DB.Where(FoodInventory{UserID: u, OffCode: o}).FirstOrCreate(&r).Error; err != nil {
		return nil, err
	}

	api := openfoodfacts.NewClient("world", "", "")
	product, err := api.Product(r.OffCode)
	if err != nil {
		return nil, err
	}

	r.Calories = product.Nutriments.EnergyKcal100G
	r.FatTotal = product.Nutriments.Fat100G
	r.FatSaturated = product.Nutriments.SaturatedFat100G
	r.FatPolyunsaturated = product.Nutriments.PolyunsaturatedFat100G
	r.FatMonounsaturated = product.Nutriments.MonounsaturatedFat100G
	r.FatTrans = product.Nutriments.TransFat100G
	r.FatCholesterol = product.Nutriments.Cholesterol100G
	r.Sodium = product.Nutriments.Sodium100G
	r.Potassium = product.Nutriments.Potassium100G
	r.Carbs = product.Nutriments.Carbohydrates100G
	r.Fibers = product.Nutriments.Fiber100G
	r.Sugars = product.Nutriments.Sugars100G
	r.Proteins = product.Nutriments.Proteins100G

	if err := DB.Save(&r).Error; err != nil {
		return nil, err
	}

	return &r, nil
}
