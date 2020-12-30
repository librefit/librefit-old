package database

import (
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/google/uuid"

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
}

// StoreImages build the mapping table between images and uploads
func StoreImages(urls *[]string, userID uint, FoodInventoryID *uint) error {
	for _, u := range *urls {
		filename, err := saveFile(u)
		if err != nil {
			return err
		}

		upload := &Upload{
			RelativeDir: "./data/img/",
			Filename:    filename,
			UserID:      userID,
		}

		if err := DB.Create(&upload).Error; err != nil {
			return err
		}

		inventoryImage := &FoodInventoryImg{
			UploadID:        upload.ID,
			FoodInventoryID: *FoodInventoryID,
		}

		if err := DB.Create(&inventoryImage).Error; err != nil {
			return err
		}
	}

	return nil
}

func saveFile(u string) (string, error) {
	filename := uuid.New().String() + filepath.Ext(path.Base(u))
	filepath := "./data/img/" + filename

	out, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	resp, err := http.Get(u)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func FindOneFoodInventory(id string) (FoodInventory, error) {
	var f FoodInventory
	err := DB.Preload("FoodInventoryImg").Find(&f, id).Error
	return f, err
}

func DeleteFoodInventory(id string) error {
	var f FoodInventory
	if err := DB.Preload("FoodInventoryImg.Upload").Preload("FoodInventoryImg").Preload("FoodDiary").First(&f, "id = ?", id).Error; err != nil {
		return err
	}

	// In case there are Foreign keys from Food Diary table
	if f.FoodDiary == nil {
		DB.Unscoped().Delete(f.FoodInventoryImg)
		return DB.Unscoped().Delete(f).Error
	}

	return DB.Delete(f).Error
}

func FindManyFoodInventory(UserID uint) ([]FoodInventory, error) {
	var f []FoodInventory
	err := DB.Preload("FoodInventoryImg.Upload").Preload("FoodInventoryImg").Find(&f, "user_id = ?", UserID).Error
	return f, err
}
