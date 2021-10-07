package controllers

import (
	"bufio"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/joeyave/statistics-project1/templates"
	"github.com/shopspring/decimal"
	"net/http"
)

var Data []decimal.Decimal

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

	buf := new(bytes.Buffer)
	buf.ReadFrom(openedFile)

	var data []decimal.Decimal

	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		float, err := decimal.NewFromString(scanner.Text())
		if err != nil {
			return
		}

		data = append(data, float)
	}

	Data = data

	c.HTML(http.StatusOK, "upload.tmpl", templates.Upload{Data: data})
}
