package controllers

import (
	"bytes"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/joeyave/statistics-project1/helpers"
	"github.com/joeyave/statistics-project1/templates"
	"net/http"
)

var Variants []*templates.Variant

func VariationalSeries(c *gin.Context) {

	x := Data

	variants := helpers.ECDF(x)
	p := helpers.PlotECDF(x)

	writer, err := p.WriterTo(400, 400, "svg")
	if err != nil {
		return
	}

	buf := new(bytes.Buffer)
	_, err = writer.WriteTo(buf)
	if err != nil {
		return
	}

	str := base64.StdEncoding.EncodeToString(buf.Bytes())

	c.HTML(http.StatusOK, "variational-series.tmpl", templates.VariationalSeries{
		Variants: variants,
		Image:    str,
	})
}
