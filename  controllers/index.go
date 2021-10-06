package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"statistics-project1/templates"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", templates.Main{
		Title: "Statistics Project 1",
	})
}
