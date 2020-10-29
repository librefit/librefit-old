package api

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	db "github.com/librefitness/librefitness/internal/database"
	serializers "github.com/librefitness/librefitness/internal/serializers"
	validators "github.com/librefitness/librefitness/internal/validators"
)

func measures(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := uint(claims["UserID"].(float64))

	measures, err := db.FindManyMeasurement(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := serializers.MeasurementsResponse(&measures)

	c.JSON(http.StatusOK, response)
}

func measureCreate(c *gin.Context) {
	measurementValidator := validators.NewMeasurementValidator()
	if err := measurementValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.SaveOne(&measurementValidator.MeasurementDb); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error (db)": err.Error()})
		return
	}

	response := serializers.MeasurementResponse(&measurementValidator.MeasurementDb)

	c.JSON(http.StatusCreated, response)
}

func measureDelete(c *gin.Context) {
	id := c.Param("id")

	if err := db.DeleteMeasurement(id); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error (db)": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func measureUpdate(c *gin.Context) {
	id := c.Param("id")

	m, err := db.FindOneMeasurement(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	measurementValidator := validators.NewMeasurementValidatorFillWith(m)
	if err := measurementValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	measurementValidator.MeasurementDb.ID = m.ID

	if err := db.SaveOne(&measurementValidator.MeasurementDb); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error (db)": err.Error()})
		return
	}

	response := serializers.MeasurementResponse(&measurementValidator.MeasurementDb)

	c.JSON(http.StatusCreated, response)
}
