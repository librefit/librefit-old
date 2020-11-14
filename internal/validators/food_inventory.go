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
		Favourite          bool    `json:"favourite"`
		Calories           float64 `json:"calories"`
		FatTotal           float64 `json:"FatTotal"`
		FatSaturated       float64 `json:"fat_saturated"`
		FatPolyunsaturated float64 `json:"fat_polyunsaturated"`
		FatMonounsaturated float64 `json:"fat_monounsaturated"`
		FatTrans           float64 `json:"fat_trans"`
		FatCholesterol     float64 `json:"fat_cholesterol"`
		Sodium             float64 `json:"sodium"`
		Potassium          float64 `json:"potassium"`
		Carbs              float64 `json:"carbs"`
		Fibers             float64 `json:"fibers"`
		Sugars             float64 `json:"sugars"`
		Proteins           float64 `json:"proteins"`
	} `json:"food_inventory"`
	FoodInventoryDb db.FoodInventory `json:"-"`
}

func NewFoodInventoryValidatorFillWith(f db.FoodInventory) FoodInventoryValidator {
	mv := NewFoodInventoryValidator()

	mv.FoodInventory.OffCode = f.OffCode
	mv.FoodInventory.ProductName = f.ProductName
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
