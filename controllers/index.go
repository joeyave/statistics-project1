package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joeyave/statistics-project1/templates"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", templates.Main{
		Title: "Statistics Project 1",
	})
}
