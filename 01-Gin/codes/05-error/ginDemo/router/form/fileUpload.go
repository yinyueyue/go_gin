package form

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func SingleFileUpload(c *gin.Context) {

	multiFile, _ := c.FormFile("file")
	sourceFile, _ := multiFile.Open()
	defer sourceFile.Close()
	filename := multiFile.Filename
	log.Printf("fileName=%v", filename)

	destFile, _ := os.OpenFile("D://"+filename, os.O_WRONLY|os.O_CREATE, 1)
	defer destFile.Close()
	_, err := io.Copy(bufio.NewWriter(destFile), bufio.NewReader(sourceFile))
	if err != nil {
		fmt.Printf("err=%v", err)
	}

	c.JSON(http.StatusOK, gin.H{"code": 2000, "msg": "ok"})
}

func MultipleFileUpload(c *gin.Context) {

	form, _ := c.MultipartForm()
	files := form.File["files"]

	for _, file := range files {
		fmt.Println(file.Filename)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))

	c.JSON(http.StatusOK, gin.H{"code": 2000, "msg": "ok"})
}

