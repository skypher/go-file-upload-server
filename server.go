package main

import (
    "fmt"
    "math"
    "net/http"
    "path/filepath"

    "github.com/gin-gonic/gin"
)

func setupHttpHandlers(router *gin.Engine) {
	router.Static("/", "./public")

	router.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		files := form.File["files"]

		for _, file := range files {
			basename := filepath.Base(file.Filename)
			filename := filepath.Join(".", "uploads", basename)
			if err := c.SaveUploadedFile(file, filename); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}

		var filenames []string
		for _, file := range files {
			filenames = append(filenames, file.Filename)
		}
		c.JSON(http.StatusOK, gin.H{"files": filenames})
	})
}

var port int16 = 3335
var maxSize int64 = int64(math.Pow10(9)) //1GB

func main() {
	router := gin.Default()
    setupHttpHandlers(router)
    fmt.Printf("Starting server on port %d", port)
    router.Run(fmt.Sprintf(":%d", port))

}
