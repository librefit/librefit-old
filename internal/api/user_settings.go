package api

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	db "github.com/librefitness/librefitness/internal/database"
	"github.com/librefitness/librefitness/internal/serializers"
	"github.com/librefitness/librefitness/internal/validators"
)

func userPreferences(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	user, err := db.FindOneUser(&db.User{Username: claims["id"].(string)})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := serializers.UserSettingResponse(&user.UserSetting)

	c.JSON(http.StatusOK, response)
}

func userPreferencesUpdate(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	user, err := db.FindOneUser(&db.User{Username: claims["id"].(string)})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	usv := validators.NewUserSettingsValidatorFillWith(user.UserSetting)
	if err := usv.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usv.UserSettingDb.ID = user.UserSetting.ID

	if err := db.SaveOne(&usv.UserSettingDb); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error (db)": err.Error()})
		return
	}

	response := serializers.UserSettingResponse(&usv.UserSettingDb)

	c.JSON(http.StatusCreated, response)
}
