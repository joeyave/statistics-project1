package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/rs/zerolog/log"
	"math/rand"
	"statistics-project1/ controllers"
)

// generate random data for line chart
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func main() {
	// Creates Gin's router.
	router := gin.New()
	router.LoadHTMLGlob("templates/*")

	router.Static("/static", "static")

	router.POST("/upload", controllers.Upload)

	router.GET("/index", controllers.Index)

	router.Any("/chart-test", func(c *gin.Context) {
		// create a new line instance
		line := charts.NewLine()
		// set some global options like Title/Legend/ToolTip or anything else
		line.SetGlobalOptions(
			charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
			charts.WithTitleOpts(opts.Title{
				Title:    "Line example in Westeros theme",
				Subtitle: "Line chart rendered by the http server this time",
			}))

		// Put data into instance
		line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
			AddSeries("Category A", generateLineItems()).
			AddSeries("Category B", generateLineItems()).
			SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

		line.Render(c.Writer)
	})

	log.Info().Msgf("Starting Gin with mode: %s", gin.Mode())
	err := router.Run(":8080")
	if err != nil {
		log.Fatal().Msgf("Error starting Gin: %v", err)
	}
}
