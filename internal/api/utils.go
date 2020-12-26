package api

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	db "github.com/librefitness/librefitness/internal/database"
)

func fileUpload(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := uint(claims["UserID"].(float64))

	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	title := c.PostForm("title")
	caption := c.PostForm("caption")
	size := c.PostForm("size")
	fID, err := strconv.ParseUint(c.PostForm("food_inventory_id"), 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	extension := filepath.Ext(file.Filename)
	uuid := uuid.New().String()
	newFileName := uuid + extension

	if err := c.SaveUploadedFile(file, "./data/img/"+newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	u := &db.Upload{
		Title:       title,
		Caption:     caption,
		Size:        size,
		RelativeDir: "./data/img/",
		Filename:    newFileName,
		UserID:      userID,
	}

	if err := db.DB.Create(&u).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error (db)": err.Error()})
		return
	}

	if fID != 0 {
		fii := &db.FoodInventoryImg{
			UploadID:        u.ID,
			FoodInventoryID: uint(fID),
		}
		if err := db.DB.Create(&fii).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error (db)": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"id": uuid})
}
