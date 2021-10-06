package controllers

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Upload(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		return
	}
	defer openedFile.Close()

	var data []float64

	scanner := bufio.NewScanner(openedFile)
	for scanner.Scan() {
		float, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return
		}

		data = append(data, float)
	}

	c.String(200, "%s", data)
}
