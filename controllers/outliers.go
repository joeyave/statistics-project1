package controllers

import (
	"bytes"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/joeyave/statistics-project1/helpers"
	"net/http"
	"strconv"
)

func Outliers(c *gin.Context) {

	data := Data

	alpha, err := strconv.ParseFloat(c.PostForm("alpha"), 64)
	if err != nil {
		alpha = 0.05
	}
	action := c.PostForm("action")
	if action == "delete-outliers" {
		data = helpers.DeleteOutliers(alpha, data)
		Data = data
	}

	outliers := helpers.Outliers(alpha, data)

	p := helpers.PlotOutliers(alpha, data)
	writerTo, err := p.WriterTo(helpers.PlotWidth, helpers.PlotHeight, "svg")
	if err != nil {
		return
	}

	buf := new(bytes.Buffer)
	writerTo.WriteTo(buf)

	str := base64.StdEncoding.EncodeToString(buf.Bytes())

	c.HTML(http.StatusOK, "outliers.tmpl", map[string]interface{}{
		"FileName": FileName,
		"Outliers": outliers,
		"Alpha":    alpha,
		"Image":    str,
	})
}
