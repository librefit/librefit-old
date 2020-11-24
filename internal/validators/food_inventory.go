package validators

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	db "github.com/librefitness/librefitness/internal/database"
)

type FoodInventoryValidator struct {
	FoodInventory struct {
		OffCode            string  `json:"off_code" binding:"required"`
		ProductName        string  `json:"product_name"`
		Description        string  `json:"description"`
		Favourite          bool    `json:"favourite"`
		Calories           float32 `json:"calories"`
		FatTotal           float32 `json:"FatTotal"`
		FatSaturated       float32 `json:"fat_saturated"`
		FatPolyunsaturated float32 `json:"fat_polyunsaturated"`
		FatMonounsaturated float32 `json:"fat_monounsaturated"`
		FatTrans           float32 `json:"fat_trans"`
		FatCholesterol     float32 `json:"fat_cholesterol"`
		Sodium             float32 `json:"sodium"`
		Potassium          float32 `json:"potassium"`
		Carbs              float32 `json:"carbs"`
		Fibers             float32 `json:"fibers"`
		Sugars             float32 `json:"sugars"`
		Proteins           float32 `json:"proteins"`
	} `json:"food_inventory"`
	FoodInventoryDb db.FoodInventory `json:"-"`
}

func NewFoodInventoryValidatorFillWith(f db.FoodInventory) FoodInventoryValidator {
	mv := NewFoodInventoryValidator()

	mv.FoodInventory.OffCode = f.OffCode
	mv.FoodInventory.ProductName = f.ProductName
	mv.FoodInventory.Description = f.Description
	mv.FoodInventory.Favourite = f.Favourite
	mv.FoodInventory.Calories = f.Calories
	mv.FoodInventory.FatTotal = f.FatTotal
	mv.FoodInventory.FatSaturated = f.FatSaturated
	mv.FoodInventory.FatPolyunsaturated = f.FatPolyunsaturated
	mv.FoodInventory.FatMonounsaturated = f.FatMonounsaturated
	mv.FoodInventory.FatTrans = f.FatTrans
	mv.FoodInventory.FatCholesterol = f.FatCholesterol
	mv.FoodInventory.Sodium = f.Sodium
	mv.FoodInventory.Potassium = f.Potassium
	mv.FoodInventory.Carbs = f.Carbs
	mv.FoodInventory.Fibers = f.Fibers
	mv.FoodInventory.Sugars = f.Sugars
	mv.FoodInventory.Proteins = f.Proteins

	return mv
}

func (self *FoodInventoryValidator) Bind(c *gin.Context) error {
	claims := jwt.ExtractClaims(c)

	if err := c.ShouldBindJSON(&self.FoodInventory); err != nil {
		return err
	}

	self.FoodInventoryDb.OffCode = self.FoodInventory.OffCode
	self.FoodInventoryDb.ProductName = self.FoodInventory.ProductName
	self.FoodInventoryDb.Description = self.FoodInventory.Description
	self.FoodInventoryDb.Favourite = self.FoodInventory.Favourite
	self.FoodInventoryDb.UserID = uint(claims["UserID"].(float64))
	self.FoodInventoryDb.Calories = self.FoodInventory.Calories
	self.FoodInventoryDb.FatTotal = self.FoodInventory.FatTotal
	self.FoodInventoryDb.FatSaturated = self.FoodInventory.FatSaturated
	self.FoodInventoryDb.FatPolyunsaturated = self.FoodInventory.FatPolyunsaturated
	self.FoodInventoryDb.FatMonounsaturated = self.FoodInventory.FatMonounsaturated
	self.FoodInventoryDb.FatTrans = self.FoodInventory.FatTrans
	self.FoodInventoryDb.FatCholesterol = self.FoodInventory.FatCholesterol
	self.FoodInventoryDb.Sodium = self.FoodInventory.Sodium
	self.FoodInventoryDb.Potassium = self.FoodInventory.Potassium
	self.FoodInventoryDb.Carbs = self.FoodInventory.Carbs
	self.FoodInventoryDb.Fibers = self.FoodInventory.Fibers
	self.FoodInventoryDb.Sugars = self.FoodInventory.Sugars
	self.FoodInventoryDb.Proteins = self.FoodInventory.Proteins

	return nil
}

func NewFoodInventoryValidator() FoodInventoryValidator {
	v := FoodInventoryValidator{}
	return v
}
