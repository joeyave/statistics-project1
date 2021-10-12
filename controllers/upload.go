package controllers

import (
	"bufio"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/joeyave/statistics-project1/global"
	"net/http"
	"strconv"
	"strings"
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

	buf := new(bytes.Buffer)
	buf.ReadFrom(openedFile)

	var data []float64

	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		text := scanner.Text()
		vals := strings.Fields(text)

		for _, val := range vals {
			float, err := strconv.ParseFloat(val, 64)
			if err != nil {
				continue
			}
			data = append(data, float)
		}
	}

	global.SetData(data)
	global.SetFileName(file.Filename)

	c.HTML(http.StatusOK, "upload.tmpl", map[string]interface{}{
		"FileName": file.Filename,
		"Data":     data,
	})
}
