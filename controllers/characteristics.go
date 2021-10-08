package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joeyave/statistics-project1/helpers"
	"github.com/joeyave/statistics-project1/templates"
	"gonum.org/v1/gonum/stat"
	"math"
	"net/http"
	"sort"
)

func Characteristics(c *gin.Context) {
	data := Data
	sort.Float64s(data)

	mean := helpers.Mean(data)

	var characteristics []*templates.Characteristic
	characteristics = append(characteristics, &templates.Characteristic{
		Name: "mean",
		Val:  mean,
	})

	median := helpers.Median(data)

	characteristics = append(characteristics, &templates.Characteristic{
		Name: "median",
		Val:  median,
	})

	variance := stat.Variance(data, nil)
	standardDeviation := math.Sqrt(variance)

	characteristics = append(characteristics, &templates.Characteristic{
		Name: "standard deviation",
		Val:  standardDeviation,
	})

	skewness := stat.Skew(data, nil)

	characteristics = append(characteristics, &templates.Characteristic{
		Name: "skewness",
		Val:  skewness,
	})

	kurtosis := helpers.Kurtosis(data)

	characteristics = append(characteristics, &templates.Characteristic{
		Name: "kurtosis",
		Val:  kurtosis,
	})

	antikurtosis := helpers.AntiKurtosis(data)

	characteristics = append(characteristics, &templates.Characteristic{
		Name: "antikurtosis",
		Val:  antikurtosis,
	})

	c.HTML(http.StatusOK, "characteristics.tmpl", templates.Characteristics{
		Characteristics: characteristics,
	})
}
