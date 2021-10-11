package controllers

import (
	"bytes"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/joeyave/statistics-project1/helpers"
	"github.com/joeyave/statistics-project1/templates"
	"net/http"
	"sort"
)

func Characteristics(c *gin.Context) {
	data := Data
	sort.Float64s(data)

	var characteristics []*templates.Characteristic

	mean := helpers.Mean(data)
	stdErr := helpers.MeanStandardError(data)
	from, to := helpers.MeanConfidenceInterval(0.05, data)

	characteristics = append(characteristics, &templates.Characteristic{
		Name:   "mean",
		Val:    mean,
		StdDev: stdErr,
		From:   from,
		To:     to,
	})

	median := helpers.Median(data)
	from, to = helpers.MedianConfidenceInterval(0.05, data)

	characteristics = append(characteristics, &templates.Characteristic{
		Name:   "median",
		Val:    median,
		StdDev: 0,
		From:   from,
		To:     to,
	})

	stdDev := helpers.StandardDeviation(data)
	stdDevStdErr := helpers.StandardDeviationStandardError(data)
	from, to = helpers.StandardDeviationConfidenceInterval(0.05, data)

	characteristics = append(characteristics, &templates.Characteristic{
		Name:   "standard deviation",
		Val:    stdDev,
		StdDev: stdDevStdErr,
		From:   from,
		To:     to,
	})

	skewness := helpers.Skewness(data)
	stdErr = helpers.SkewnessStandardError(data)
	from, to = helpers.SkewnessConfidenceInterval(0.05, data)

	characteristics = append(characteristics, &templates.Characteristic{
		Name:   "skewness",
		Val:    skewness,
		StdDev: stdErr,
		From:   from,
		To:     to,
	})

	kurtosis := helpers.Kurtosis(data)
	stdErr = helpers.KurtosisStandardError(data)
	from, to = helpers.KurtosisConfidenceInterval(0.05, data)

	characteristics = append(characteristics, &templates.Characteristic{
		Name:   "kurtosis",
		Val:    kurtosis,
		StdDev: stdErr,
		From:   from,
		To:     to,
	})

	antikurtosis := helpers.AntiKurtosis(data)

	characteristics = append(characteristics, &templates.Characteristic{
		Name: "antikurtosis",
		Val:  antikurtosis,
	})

	characteristics = append(characteristics, &templates.Characteristic{
		Name: "min",
		Val:  helpers.Min(data),
	})

	characteristics = append(characteristics, &templates.Characteristic{
		Name: "max",
		Val:  helpers.Max(data),
	})

	p := helpers.PlotPDF(data)

	writerTo, err := p.WriterTo(400, 400, "svg")
	if err != nil {
		return
	}

	buf := new(bytes.Buffer)
	writerTo.WriteTo(buf)

	str := base64.StdEncoding.EncodeToString(buf.Bytes())

	c.HTML(http.StatusOK, "characteristics.tmpl", templates.Characteristics{
		Characteristics: characteristics,
		Image:           str,
	})
}
