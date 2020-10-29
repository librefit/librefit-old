package api

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	db "github.com/librefitness/librefitness/internal/database"
	"github.com/librefitness/librefitness/internal/serializers"
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

func userPreferences(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	user, err := db.FindOneUser(&db.User{Username: claims["id"].(string)})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// err = user.AddJSON()
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	// 	return
	// }
	spew.Dump(user)

	c.JSON(http.StatusOK, user)
}

func userPreferencesUpdate(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	user, err := db.FindOneUser(&db.User{Username: claims["id"].(string)})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = user.AddJSON()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
