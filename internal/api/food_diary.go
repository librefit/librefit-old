package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/librefitness/librefitness/internal/database"
	serializers "github.com/librefitness/librefitness/internal/serializers"
	"github.com/librefitness/librefitness/internal/validators"
)

func foodDiary(c *gin.Context) {
	c.JSON(http.StatusOK, "ale")
}

func foodDiaryCreate(c *gin.Context) {
	fdv := validators.NewFoodDiaryValidator()
	if err := fdv.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.SaveOne(&fdv.FoodDiaryDb); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error (db)": err.Error()})
		return
	}

	response := serializers.FoodDiaryResponse(&fdv.FoodDiaryDb)

	c.JSON(http.StatusCreated, response)
}

func foodDiaryDelete(c *gin.Context) {
	c.JSON(http.StatusOK, "ale")
}

func foodDiaryUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, "ale")
}
