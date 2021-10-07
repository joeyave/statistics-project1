package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joeyave/statistics-project1/templates"
	"github.com/shopspring/decimal"
	"net/http"
	"sort"
)

var Variants []*templates.Variant

func VariationalSeries(c *gin.Context) {

	variantToNumMap := make(map[decimal.Decimal]int)

	data := Data
	sort.Slice(data, func(i, j int) bool {
		return data[i].LessThan(data[j])
	})

	for _, v := range Data {
		_, exists := variantToNumMap[v]
		if exists {
			variantToNumMap[v] += 1
		} else {
			variantToNumMap[v] = 1
		}
	}

	keys := make([]decimal.Decimal, 0, len(variantToNumMap))
	for k := range variantToNumMap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].LessThan(keys[j])
	})

	var variants []*templates.Variant

	for _, v := range keys {
		num := variantToNumMap[v]
		variant := templates.Variant{
			X: v,
			N: num,
			P: decimal.NewFromInt32(int32(num)).Div(decimal.NewFromInt32(int32(len(Data)))),
		}

		variants = append(variants, &variant)
	}

	for i, v := range variants {

		if i == 0 {
			v.F = v.P
		} else {
			// v.F += v.P + variants[i-1].F

			v.F = variants[i-1].F.Add(v.P)
		}
	}

	Variants = variants

	c.HTML(http.StatusOK, "variational-series.tmpl", templates.VariationalSeries{
		Variants: variants,
	})
}
