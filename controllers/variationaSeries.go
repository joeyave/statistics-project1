package controllers

import (
	"bytes"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/joeyave/statistics-project1/templates"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"net/http"
	"sort"
)

var Variants []*templates.Variant

func VariationalSeries(c *gin.Context) {

	variantToNumMap := make(map[float64]int)

	data := Data
	sort.Float64s(data)

	for _, v := range Data {
		_, exists := variantToNumMap[v]
		if exists {
			variantToNumMap[v] += 1
		} else {
			variantToNumMap[v] = 1
		}
	}

	keys := make([]float64, 0, len(variantToNumMap))
	for k := range variantToNumMap {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	var variants []*templates.Variant

	for _, v := range keys {
		num := variantToNumMap[v]
		variant := templates.Variant{
			X: v,
			N: num,
			P: float64(num) / float64(len(Data)),
		}

		variants = append(variants, &variant)
	}

	for i, v := range variants {

		if i == 0 {
			v.F = v.P
		} else {
			v.F = variants[i-1].F + v.P
		}
	}

	Variants = variants

	p := plot.New()
	p.Add(plotter.NewGrid())

	dot1 := plotter.XY{X: Variants[0].X - 1, Y: 0}
	dot2 := plotter.XY{X: Variants[0].X, Y: 0}
	err := plotutil.AddLines(p, plotter.XYs{dot1, dot2})
	if err != nil {
		return
	}

	for i := 0; i < len(Variants); i++ {
		dot1 := plotter.XY{X: Variants[i].X, Y: Variants[i].F}

		dot2 := plotter.XY{}
		if i == len(Variants)-1 {
			dot2 = plotter.XY{X: Variants[len(Variants)-1].X + 1, Y: 1}
		} else {
			dot2 = plotter.XY{X: Variants[i+1].X, Y: Variants[i].F}
		}

		err := plotutil.AddLines(p, plotter.XYs{dot1, dot2})
		if err != nil {
			return
		}

		err = plotutil.AddScatters(p, plotter.XYs{dot1})
		if err != nil {
			return
		}
	}

	to, err := p.WriterTo(4*vg.Inch, 4*vg.Inch, "svg")
	if err != nil {
		return
	}

	buf := new(bytes.Buffer)
	to.WriteTo(buf)

	str := base64.StdEncoding.EncodeToString(buf.Bytes())

	c.HTML(http.StatusOK, "variational-series.tmpl", templates.VariationalSeries{
		Variants: variants,
		Image:    str,
	})
}
