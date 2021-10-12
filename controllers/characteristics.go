package controllers

import (
	"bytes"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/joeyave/statistics-project1/global"
	"github.com/joeyave/statistics-project1/helpers"
	"github.com/joeyave/statistics-project1/templates"
	"net/http"
	"sort"
)

func Characteristics(c *gin.Context) {
	x := global.DataCopy()
	sort.Float64s(x)

	var characteristics []*templates.Characteristic

	mean := helpers.Mean(x)
	stdErr := helpers.MeanStandardError(x)
	from, to := helpers.MeanConfidenceInterval(0.05, x)

	characteristics = append(characteristics, &templates.Characteristic{
		Name:   "mean",
		Val:    mean,
		StdDev: stdErr,
		From:   from,
		To:     to,
	})

	median := helpers.Median(x)
	from, to = helpers.MedianConfidenceInterval(0.05, x)

	characteristics = append(characteristics, &templates.Characteristic{
		Name:   "median",
		Val:    median,
		StdDev: 0,
		From:   from,
		To:     to,
	})

	stdDev := helpers.StandardDeviation(x)
	stdDevStdErr := helpers.StandardDeviationStandardError(x)
	from, to = helpers.StandardDeviationConfidenceInterval(0.05, x)

	characteristics = append(characteristics, &templates.Characteristic{
		Name:   "standard deviation",
		Val:    stdDev,
		StdDev: stdDevStdErr,
		From:   from,
		To:     to,
	})

	skewness := helpers.Skewness(x)
	stdErr = helpers.SkewnessStandardError(x)
	from, to = helpers.SkewnessConfidenceInterval(0.05, x)

	characteristics = append(characteristics, &templates.Characteristic{
		Name:   "skewness",
		Val:    skewness,
		StdDev: stdErr,
		From:   from,
		To:     to,
	})

	kurtosis := helpers.Kurtosis(x)
	stdErr = helpers.KurtosisStandardError(x)
	from, to = helpers.KurtosisConfidenceInterval(0.05, x)

	characteristics = append(characteristics, &templates.Characteristic{
		Name:   "kurtosis",
		Val:    kurtosis,
		StdDev: stdErr,
		From:   from,
		To:     to,
	})

	antikurtosis := helpers.AntiKurtosis(x)

	characteristics = append(characteristics, &templates.Characteristic{
		Name: "antikurtosis",
		Val:  antikurtosis,
	})

	characteristics = append(characteristics, &templates.Characteristic{
		Name: "min",
		Val:  helpers.Min(x),
	})

	characteristics = append(characteristics, &templates.Characteristic{
		Name: "max",
		Val:  helpers.Max(x),
	})

	p := helpers.PlotNormalPDF(x)

	writerTo, err := p.WriterTo(helpers.PlotWidth, helpers.PlotHeight, "svg")
	if err != nil {
		return
	}

	buf := new(bytes.Buffer)
	writerTo.WriteTo(buf)

	str := base64.StdEncoding.EncodeToString(buf.Bytes())

	c.HTML(http.StatusOK, "characteristics.tmpl", map[string]interface{}{
		"FileName":        global.FileName(),
		"Characteristics": characteristics,
		"Image":           str,
	})
}
