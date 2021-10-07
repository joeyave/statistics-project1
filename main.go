package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joeyave/statistics-project1/controllers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"html/template"
	"os"
	"time"
)

func main() {
	out := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
	log.Logger = zerolog.New(out).Level(zerolog.GlobalLevel()).With().Timestamp().Logger()

	router := gin.New()
	router.SetFuncMap(template.FuncMap{
		"sub": func(i, j int) int {
			return i - j
		},
	})

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "static")

	router.GET("/index", controllers.Index)
	router.POST("/upload", controllers.Upload)

	router.GET("/variationalSeries", controllers.VariationalSeries)
	router.GET("/classes", controllers.ClassesGet)
	router.POST("/classes", controllers.ClassesPost)

	log.Info().Msgf("Starting Gin with mode: %s", gin.Mode())
	err := router.Run(":8080")
	if err != nil {
		log.Fatal().Msgf("Error starting Gin: %v", err)
	}
}
