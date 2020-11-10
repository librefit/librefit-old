package api

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/librefitness/librefitness/internal/database"
	db "github.com/librefitness/librefitness/internal/database"
	serializers "github.com/librefitness/librefitness/internal/serializers"
	"github.com/librefitness/librefitness/internal/validators"
)

func foodDiary(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := uint(claims["UserID"].(float64))

	f, err := db.FindManyFoodDiary(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := serializers.FoodDiariesResponse(&f)

	c.JSON(http.StatusOK, response)
}

func foodDiaryCreate(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := uint(claims["UserID"].(float64))

	fdv := validators.NewFoodDiaryValidator()
	if err := fdv.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var fi database.FoodInventory
	if err := db.DB.FirstOrCreate(&fi, database.FoodInventory{UserID: userID, OffCode: fdv.FoodDiary.OffCode}).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error (db)": err.Error()})
		return
	}

	fdv.FoodDiaryDb.FoodInventoryID = fi.ID

	if err := db.SaveOne(&fdv.FoodDiaryDb); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error (db)": err.Error()})
		return
	}

	response := serializers.FoodDiaryResponse(&fdv.FoodDiaryDb)

	c.JSON(http.StatusCreated, response)
}

func foodDiaryDelete(c *gin.Context) {
	id := c.Param("id")

	if err := db.DeleteFoodDiary(id); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error (db)": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func foodDiaryUpdate(c *gin.Context) {
	id := c.Param("id")

	f, err := db.FindOneFoodDiary(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	fdv := validators.NewFoodDiaryValidatorFillWith(f)
	if err := fdv.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fdv.FoodDiaryDb.ID = f.ID

	if err := db.SaveOne(&fdv.FoodDiaryDb); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error (db)": err.Error()})
		return
	}

	response := serializers.FoodDiaryResponse(&fdv.FoodDiaryDb)

	c.JSON(http.StatusCreated, response)
}
