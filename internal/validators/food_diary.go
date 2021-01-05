package validators

import (
	"fmt"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	db "github.com/librefitness/librefitness/internal/database"
)

type FoodDiaryValidator struct {
	FoodDiary struct {
		MealTypeID      string  `json:"meal_type" binding:"required"`
		Quantity        float32 `json:"quantity" binding:"required"`
		QuantityUnitID  uint    `json:"quantity_unit" binding:"required"`
		Comments        string  `json:"comments"`
		Date            string  `json:"date" binding:"required"`
		FoodInventoryID uint    `json:"food_inventory_id" binding:"required"`
	} `json:"food_diary"`
	FoodDiaryDb db.FoodDiary `json:"-"`
}

func NewFoodDiaryValidatorFillWith(f db.FoodDiary) FoodDiaryValidator {
	mv := NewFoodDiaryValidator()

	mv.FoodDiary.MealTypeID = f.MealTypeID
	mv.FoodDiary.Quantity = f.Quantity
	mv.FoodDiary.QuantityUnitID = f.QuantityUnitID
	mv.FoodDiary.Comments = f.Comments
	mv.FoodDiary.Date = fmt.Sprintf("%s", f.Date)
	mv.FoodDiary.FoodInventoryID = f.FoodInventoryID

	return mv
}

func (self *FoodDiaryValidator) Bind(c *gin.Context) error {
	claims := jwt.ExtractClaims(c)

	if err := c.ShouldBindJSON(&self.FoodDiary); err != nil {
		return err
	}

	date, err := time.Parse("2006-01-02", self.FoodDiary.Date)
	if err != nil {
		return err
	}

	self.FoodDiaryDb.MealTypeID = self.FoodDiary.MealTypeID
	self.FoodDiaryDb.Quantity = self.FoodDiary.Quantity
	self.FoodDiaryDb.QuantityUnitID = self.FoodDiary.QuantityUnitID
	self.FoodDiaryDb.FoodInventoryID = self.FoodDiary.FoodInventoryID
	self.FoodDiaryDb.Comments = self.FoodDiary.Comments
	self.FoodDiaryDb.Date = date
	self.FoodDiaryDb.UserID = uint(claims["UserID"].(float64))

	return nil
}

func NewFoodDiaryValidator() FoodDiaryValidator {
	return FoodDiaryValidator{}
}
