package controllers

import (
	"bytes"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/joeyave/statistics-project1/global"
	"github.com/joeyave/statistics-project1/helpers"
	"math"
	"net/http"
	"strconv"
)

func Classes(c *gin.Context) {

	x := global.DataCopy()

	h, err := strconv.ParseFloat(c.PostForm("h"), 64)
	if err != nil {
		h = 0
	}

	M, err := strconv.Atoi(c.PostForm("M"))
	if err != nil {
		M = int(1 + 1.44*math.Log(float64(len(x))))
	}

	classes := helpers.Classes(M, x)

	if h == 0 {
		h = helpers.Scott(x)
	}

	p, _ := helpers.PlotHistogram(M, h, x)

	to, err := p.WriterTo(helpers.PlotWidth, helpers.PlotHeight, "svg")
	if err != nil {
		return
	}

	buf := new(bytes.Buffer)
	to.WriteTo(buf)

	str := base64.StdEncoding.EncodeToString(buf.Bytes())

	c.HTML(http.StatusOK, "classes.tmpl", map[string]interface{}{
		"FileName": global.FileName(),
		"Classes":  classes,
		"Image":    str,
		"M":        M,
		"H":        h,
	})
}
