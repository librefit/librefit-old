package api

import (
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	db "github.com/librefitness/librefitness/internal/database"
	"github.com/librefitness/librefitness/internal/serializers"
)

func statsFluids(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := uint(claims["UserID"].(float64))

	sc := db.FluidSearch{UserID: userID}

	fluids, err := db.FindFluids(&sc)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	res, err := serializers.StatsTotalFluidsPerDay(&fluids)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func statsFoodDiary(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := uint(claims["UserID"].(float64))

	start, err := time.Parse("2006-01-02", c.DefaultQuery("start", "1900-01-01")) // Just a old-random date
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	end, err := time.Parse("2006-01-02", c.DefaultQuery("end", time.Now().Format("2006-01-02")))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	f, err := db.FindManyFoodDiary(userID, start, end)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := serializers.FoodDiaryStatsResponse(&f)

	c.JSON(http.StatusOK, response)
}
