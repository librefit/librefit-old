package api

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	db "github.com/librefitness/librefitness/internal/database"
	"github.com/librefitness/librefitness/internal/serializers"
	validators "github.com/librefitness/librefitness/internal/validators"
)

func fluids(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := uint(claims["UserID"].(float64))

	sc := db.FluidSearch{UserID: userID}

	fluids, err := db.FindFluids(&sc)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := serializers.FluidsResponse(&fluids)

	c.JSON(http.StatusOK, response)
}

func fluidsWithRange(c *gin.Context) {
	fs := db.FluidSearch{}

	fluidSearchValidator := validators.FluidSearchValidator{}
	err := fluidSearchValidator.Bind(c, &fs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fluids, err := db.FindFluids(&fs)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := serializers.FluidsResponse(&fluids)

	c.JSON(http.StatusOK, response)
}

func fluidCreate(c *gin.Context) {
	fluidValidator := validators.NewFluidValidator()
	if err := fluidValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.SaveOne(&fluidValidator.FluidDb); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error (db)": err.Error()})
		return
	}

	response := serializers.FluidResponse(&fluidValidator.FluidDb)

	c.JSON(http.StatusCreated, response)
}

func fluidDelete(c *gin.Context) {
	id := c.Param("id")

	if err := db.DeleteFluid(id); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error (db)": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func fluidUpdate(c *gin.Context) {
	id := c.Param("id")

	f, err := db.FindOneFluid(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	fluidValidator := validators.NewFluidValidatorFillWith(&f)
	if err := fluidValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fluidValidator.FluidDb.ID = f.ID

	if err := db.SaveOne(&fluidValidator.FluidDb); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error (db)": err.Error()})
		return
	}

	response := serializers.FluidResponse(&fluidValidator.FluidDb)

	c.JSON(http.StatusCreated, response)
}
