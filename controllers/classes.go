package controllers

import (
	"bytes"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/joeyave/statistics-project1/templates"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"math"
	"net/http"
	"sort"
	"strconv"
)

var Classes []*templates.Class

func ClassesGet(c *gin.Context) {
	c.HTML(http.StatusOK, "classes.tmpl", nil)
}

func ClassesPost(c *gin.Context) {

	number := c.PostForm("number")

	M, err := strconv.Atoi(number)
	if err != nil {
		return
	}

	xMin := Variants[0].X
	xMax := Variants[len(Variants)-1].X

	h := (xMax - xMin) / float64(M)

	var classes []*templates.Class

	for i := 0; i < M; i++ {
		class := templates.Class{
			XFrom: xMin + (h * float64(i)),
		}
		class.XTo = class.XFrom + h

		classes = append(classes, &class)
	}

	for i := range classes {
		for _, v := range Variants {
			if i == len(classes)-1 {
				if v.X >= classes[i].XFrom && v.X <= classes[i].XTo {
					classes[i].N++
				}
			} else {
				if v.X >= classes[i].XFrom && v.X < classes[i].XTo {
					classes[i].N++
				}
			}
		}
	}

	for i := range classes {
		classes[i].P = float64(classes[i].N) / float64(len(Variants))

		if i == 0 {
			classes[i].F = classes[i].P
		} else {
			classes[i].F = classes[i-1].F + classes[i].P
		}
	}

	Classes = classes

	p := plot.New()

	var variants plotter.XYs

	for _, v := range Variants {
		xy := plotter.XY{X: v.X, Y: v.P}
		variants = append(variants, xy)
	}

	histogram, err := plotter.NewHistogram(variants, M)
	if err != nil {
		return
	}

	p.Add(histogram)

	kde := plotter.NewFunction(func(x float64) float64 {

		data := Data
		sort.Float64s(data)
		variance := stat.Variance(data, nil)
		stdDev := math.Sqrt(variance)

		h = stdDev * math.Pow(float64(len(data)), -0.2)

		kSum := 0.
		for _, val := range Data {
			u := (x - val) / h
			k := 1 / math.Sqrt(2*math.Pi) * math.Exp(-(math.Pow(u, 2) / 2))
			kSum += k
		}

		y := (1 / (float64(len(Data)) * h)) * kSum

		return y
	})

	kde.Color = color.RGBA{R: 255, A: 255}
	kde.Width = vg.Points(2)
	p.Add(kde)

	to, err := p.WriterTo(4*vg.Inch, 4*vg.Inch, "svg")
	if err != nil {
		return
	}

	buf := new(bytes.Buffer)
	to.WriteTo(buf)

	str := base64.StdEncoding.EncodeToString(buf.Bytes())

	c.HTML(http.StatusOK, "classes.tmpl", templates.Classes{
		Classes: classes,
		Image:   str,
	})
}
