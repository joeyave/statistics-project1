package controllers

import (
	"bytes"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/joeyave/statistics-project1/helpers"
	"net/http"
)

func EmpiricalDistribution(c *gin.Context) {

	x := Data

	variants := helpers.Variants(helpers.EmpiricalCDF(x), x)
	p := helpers.PlotEmpiricalCDF(x)

	writer, err := p.WriterTo(helpers.PlotWidth, helpers.PlotHeight, "svg")
	if err != nil {
		return
	}

	buf := new(bytes.Buffer)
	_, err = writer.WriteTo(buf)
	if err != nil {
		return
	}

	str := base64.StdEncoding.EncodeToString(buf.Bytes())

	c.HTML(http.StatusOK, "distribution.tmpl", map[string]interface{}{
		"Variants": variants,
		"FileName": FileName,
		"Image":    str,
	})
}
