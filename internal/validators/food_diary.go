package validators

import (
	"fmt"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	db "github.com/librefitness/librefitness/internal/database"
)

type FoodDiaryValidator struct {
	FoodDiary struct {
		MealType        string `json:"meal_type" binding:"required"`
		Quantity        string `json:"quantity" binding:"required"`
		QuantityUnit    string `json:"quantity_unit" binding:"required"`
		Comments        string `json:"comments"`
		Date            string `json:"date" binding:"required"`
		FoodInventoryID uint   `json:"food_inventory_id" binding:"required"`
	} `json:"food_diary"`
	FoodDiaryDb db.FoodDiary `json:"-"`
}

func NewFoodDiaryValidatorFillWith(f db.FoodDiary) FoodDiaryValidator {
	mv := NewFoodDiaryValidator()

	mv.FoodDiary.MealType = f.MealType
	mv.FoodDiary.Quantity = fmt.Sprintf("%f", f.Quantity)
	mv.FoodDiary.QuantityUnit = f.QuantityUnit
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

	quantity, err := strconv.ParseFloat(self.FoodDiary.Quantity, 32)
	if err != nil {
		return err
	}

	self.FoodDiaryDb.MealType = self.FoodDiary.MealType
	self.FoodDiaryDb.Quantity = float32(quantity)
	self.FoodDiaryDb.QuantityUnit = self.FoodDiary.QuantityUnit
	self.FoodDiaryDb.Comments = self.FoodDiary.Comments
	self.FoodDiaryDb.Date = date
	self.FoodDiaryDb.FoodInventoryID = self.FoodDiary.FoodInventoryID
	self.FoodDiaryDb.UserID = uint(claims["UserID"].(float64))

	return nil
}

func NewFoodDiaryValidator() FoodDiaryValidator {
	v := FoodDiaryValidator{}
	return v
}
