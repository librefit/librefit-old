package api

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	db "github.com/librefitness/librefitness/internal/database"
	"github.com/librefitness/librefitness/internal/serializers"
	"github.com/librefitness/librefitness/internal/validators"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func userInfo(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	user, err := db.FindOneUser(&db.User{Username: claims["id"].(string)})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := serializers.UserResponse(&user)

	c.JSON(http.StatusOK, response)
}

func updateUser(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	user, err := db.FindOneUser(&db.User{Username: claims["id"].(string)})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	userValidator := validators.NewUserValidatorFillWith(user)
	if err := userValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userValidator.UserDb.ID = user.ID

	if err := db.SaveOne(&userValidator.UserDb); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error (db)": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
