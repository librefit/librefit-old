package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

func fileUpload(c *gin.Context) {
	file, _ := c.FormFile("file")
	paths := []string{
		"./data/img",
		"./data/img/user_profile",
		"./data/img/food_inventory",
	}

	location := c.PostForm("location")
	name := c.PostForm("name")

	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 0700)
		}
	}

	err := c.SaveUploadedFile(file, fmt.Sprintf("./data/img/%s/%s", location, name))
	if err != nil {
		spew.Dump(err)
	}
	c.String(http.StatusOK, fmt.Sprintf("%s uploaded into data/img/%s/%s uploaded!", name, location, name))
}
