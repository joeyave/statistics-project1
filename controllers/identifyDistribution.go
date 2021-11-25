package controllers

import (
	"bytes"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/joeyave/statistics-project1/global"
	"github.com/joeyave/statistics-project1/helpers"
	"github.com/joeyave/statistics-project1/templates"
	"gonum.org/v1/plot/plotter"
	"image/color"
	"math"
	"net/http"
)

func IdentifyDistribution(c *gin.Context) {

	x := global.DataCopy()

	mle := helpers.RayleighMLE(x)
	stdDev := helpers.RayleighMLEStandardDeviation(x)
	low, high := helpers.RayleighMLEConfidenceInterval(0.05, x)
	characteristic := templates.Characteristic{
		Name:   "sigma",
		Val:    mle,
		StdDev: stdDev,
		From:   low,
		To:     high,
	}

	variants := helpers.Variants(helpers.EmpiricalCDF(x), x)

	p := helpers.DistributionIdentificationPlot(func(x_i float64) float64 {
		y := math.Sqrt(2 * math.Log(1/(1-helpers.EmpiricalCDF(x)(x_i))))
		return y
	}, x)

	line := plotter.NewFunction(func(x_i float64) float64 {
		y := math.Sqrt(2 * math.Log(1/(1-helpers.RayleighCDF(mle)(x_i))))
		return (y)
	})
	line.Color = color.RGBA{G: 255, A: 255}

	p.Add(line)

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

	h := helpers.Scott(x)
	M := int(1 + 1.44*math.Log(float64(len(x))))

	p, width := helpers.PlotHistogram(M, h, x)
	pdf := plotter.NewFunction(func(x_i float64) float64 {
		return helpers.RayleighPDF(mle)(x_i) * width
	})
	pdf.Color = color.RGBA{G: 255, A: 255}
	pdf.Width = 2
	p.Add(pdf)

	writer, err = p.WriterTo(helpers.PlotWidth, helpers.PlotHeight, "svg")
	if err != nil {
		return
	}

	buf = new(bytes.Buffer)
	_, err = writer.WriteTo(buf)
	if err != nil {
		return
	}

	str2 := base64.StdEncoding.EncodeToString(buf.Bytes())

	p = helpers.PlotEmpiricalCDF(x)
	cdf := plotter.NewFunction(helpers.RayleighCDF(mle))
	cdf.Color = color.RGBA{G: 255, A: 255}
	cdf.Width = 2
	p.Add(cdf)

	writer, err = p.WriterTo(helpers.PlotWidth, helpers.PlotHeight, "svg")
	if err != nil {
		return
	}

	buf = new(bytes.Buffer)
	_, err = writer.WriteTo(buf)
	if err != nil {
		return
	}

	str3 := base64.StdEncoding.EncodeToString(buf.Bytes())

	z := helpers.KolmogorovZ(x, mle)

	k := helpers.KolmogorovFunction(z, x)
	// quantileK := helpers.QuantileK(1 - 0.05)

	c.HTML(http.StatusOK, "distribution.tmpl", map[string]interface{}{
		"Variants":          variants,
		"FileName":          global.FileName(),
		"Image":             str,
		"HistogramImage":    str2,
		"EmpiricalCDFImage": str3,
		"Characteristics":   []*templates.Characteristic{&characteristic},
		"Z":                 z,
		"P":                 1 - k,
		"Alpha":             0.05,
		"QuantileK":         "?",
	})
}
