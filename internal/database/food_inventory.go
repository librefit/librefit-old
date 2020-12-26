package database

import (
	"github.com/davecgh/go-spew/spew"
	openfoodfacts "github.com/mogaal/openfoodfacts-go"
	"gorm.io/gorm"
)

type FoodInventory struct {
	gorm.Model
	OffCode          string
	ProductName      string
	Description      string
	Favourite        bool
	UserID           uint
	FoodDiary        []FoodDiary
	FoodInventoryImg []FoodInventoryImg

	// Important values to save so we're not always querying OFF
	Calories           float32
	FatTotal           float32
	FatSaturated       float32
	FatPolyunsaturated float32
	FatMonounsaturated float32
	FatTrans           float32
	FatCholesterol     float32
	Sodium             float32
	Potassium          float32
	Carbs              float32
	Fibers             float32
	Sugars             float32
	Proteins           float32
}

type FoodInventoryImg struct {
	gorm.Model
	UploadID        uint
	FoodInventoryID uint
	Upload          Upload
	FoodInventory   FoodInventory
}

func FindOneFoodInventory(id string) (FoodInventory, error) {
	var f FoodInventory
	err := DB.Preload("FoodInventoryImg").Find(&f, id).Error
	return f, err
}

func DeleteFoodInventory(id string) error {
	var f FoodInventory
	if err := DB.Preload("FoodDiary").First(&f, "id = ?", id).Error; err != nil {
		return err
	}

	// In case there are Foreign keys from Food Diary table
	if f.FoodDiary == nil {
		return DB.Unscoped().Delete(f).Error
	}

	return DB.Delete(f).Error
}

func FindManyFoodInventory(UserID uint) ([]FoodInventory, error) {
	var f []FoodInventory
	err := DB.Preload("FoodInventoryImg").Find(&f, "user_id = ?", UserID).Error

	spew.Dump(f)
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

	r.ProductName = product.ProductName
	r.Calories = float32(product.Nutriments.EnergyKcal100G)
	r.FatTotal = float32(product.Nutriments.Fat100G)
	r.FatSaturated = float32(product.Nutriments.SaturatedFat100G)
	r.FatPolyunsaturated = float32(product.Nutriments.PolyunsaturatedFat100G)
	r.FatMonounsaturated = float32(product.Nutriments.MonounsaturatedFat100G)
	r.FatTrans = float32(product.Nutriments.TransFat100G)
	r.FatCholesterol = float32(product.Nutriments.Cholesterol100G)
	r.Sodium = float32(product.Nutriments.Sodium100G)
	r.Potassium = float32(product.Nutriments.Potassium100G)
	r.Carbs = float32(product.Nutriments.Carbohydrates100G)
	r.Fibers = float32(product.Nutriments.Fiber100G)
	r.Sugars = float32(product.Nutriments.Sugars100G)
	r.Proteins = float32(product.Nutriments.Proteins100G)

	if err := DB.Save(&r).Error; err != nil {
		return nil, err
	}

	return &r, nil
}
