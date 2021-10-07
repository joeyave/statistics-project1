package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/shopspring/decimal"
	"math/rand"
)

func VariationalSeriesChart(c *gin.Context) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: "Variational series chart",
		}))

	var variantsX []decimal.Decimal
	for _, v := range Variants {
		variantsX = append(variantsX, v.X)
	}

	line.SetXAxis(variantsX).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: false}))

	for i := 0; i < len(Variants)-1; i++ {
		items := make([]opts.LineData, 0)

		items = append(items, opts.LineData{Value: []interface{}{i, Variants[i].F}}, opts.LineData{Value: []interface{}{i + 1, Variants[i].F}, Symbol: "none"})
		line.AddSeries("", items)
	}

	marshal, err := json.Marshal(line.JSON())
	if err != nil {
		return
	}

	fmt.Println(string(marshal))
	// Put data into instance
	// line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
	// 	AddSeries("Category A", generateLineItems()).
	// 	AddSeries("Category B", generateLineItems()).
	// 	SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	line.Render(c.Writer)
}

// generate random data for line chart
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(100)})
	}
	return items
}
