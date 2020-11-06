package api

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	db "github.com/librefitness/librefitness/internal/database"
	serializers "github.com/librefitness/librefitness/internal/serializers"
	validators "github.com/librefitness/librefitness/internal/validators"
)

func foodInventory(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := uint(claims["UserID"].(float64))

	f, err := db.FindManyFoodInventory(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := serializers.FoodInventoriesResponse(&f)

	c.JSON(http.StatusOK, response)
}

func foodInventoryCreate(c *gin.Context) {
	fiv := validators.NewFoodInventoryValidator()
	if err := fiv.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.SaveOne(&fiv.FoodInventoryDb); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error (db)": err.Error()})
		return
	}

	response := serializers.FoodInventoryResponse(&fiv.FoodInventoryDb)

	c.JSON(http.StatusCreated, response)
}

func foodInventoryDelete(c *gin.Context) {
	id := c.Param("id")

	if err := db.DeleteFoodInventory(id); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error (db)": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func foodInventoryUpdate(c *gin.Context) {
	id := c.Param("id")

	f, err := db.FindOneFoodInventory(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	fiv := validators.NewFoodInventoryValidatorFillWith(f)
	if err := fiv.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fiv.FoodInventoryDb.ID = f.ID

	if err := db.SaveOne(&fiv.FoodInventoryDb); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error (db)": err.Error()})
		return
	}

	response := serializers.FoodInventoryResponse(&fiv.FoodInventoryDb)

	c.JSON(http.StatusCreated, response)
}
