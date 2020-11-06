package validators

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	db "github.com/librefitness/librefitness/internal/database"
)

type FoodInventoryValidator struct {
	FoodInventory struct {
		OffCode   string `json:"off_code" binding:"required"`
		Favourite bool   `json:"favourite"`
	} `json:"food_inventory"`
	FoodInventoryDb db.FoodInventory `json:"-"`
}

func NewFoodInventoryValidatorFillWith(f db.FoodInventory) FoodInventoryValidator {
	mv := NewFoodInventoryValidator()
	mv.FoodInventory.OffCode = f.OffCode
	mv.FoodInventory.Favourite = f.Favourite
	return mv
}

func (self *FoodInventoryValidator) Bind(c *gin.Context) error {
	claims := jwt.ExtractClaims(c)

	if err := c.ShouldBindJSON(&self.FoodInventory); err != nil {
		return err
	}

	self.FoodInventoryDb.OffCode = self.FoodInventory.OffCode
	self.FoodInventoryDb.Favourite = self.FoodInventory.Favourite
	self.FoodInventoryDb.UserID = uint(claims["UserID"].(float64))

	return nil
}

func NewFoodInventoryValidator() FoodInventoryValidator {
	v := FoodInventoryValidator{}
	return v
}
