package controllers

import (
	"bytes"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/joeyave/statistics-project1/helpers"
	"net/http"
	"strconv"
)

func ClassesGet(c *gin.Context) {
	c.HTML(http.StatusOK, "classes.tmpl", nil)
}

func ClassesPost(c *gin.Context) {

	x := Data

	h, err := strconv.ParseFloat(c.PostForm("h"), 64)
	if err != nil {
		h = 0
	}

	M, err := strconv.Atoi(c.PostForm("M"))
	if err != nil {
		return
	}

	classes := helpers.Classes(M, x)

	if h == 0 {
		h = helpers.Scott(x)
	}

	p := helpers.PlotHistogram(M, h, x)

	to, err := p.WriterTo(helpers.PlotWidth, helpers.PlotHeight, "svg")
	if err != nil {
		return
	}

	buf := new(bytes.Buffer)
	to.WriteTo(buf)

	str := base64.StdEncoding.EncodeToString(buf.Bytes())

	c.HTML(http.StatusOK, "classes.tmpl", map[string]interface{}{
		"FileName": FileName,
		"Classes":  classes,
		"Image":    str,
		"M":        M,
		"H":        h,
	})
}
