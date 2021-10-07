package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joeyave/statistics-project1/templates"
	"github.com/shopspring/decimal"
	"net/http"
	"strconv"
)

func Classes(c *gin.Context) {

	c.HTML(http.StatusOK, "classes.tmpl", nil)
}

func ClassesPost(c *gin.Context) {

	number := c.PostForm("number")

	M, err := strconv.Atoi(number)
	if err != nil {
		return
	}

	xMin := Variants[0].X
	xMax := Variants[len(Variants)-1].X

	h := xMax.Sub(xMin).Div(decimal.NewFromInt32(int32(M)))

	var classes []*templates.Class

	for i := 0; i < M; i++ {
		class := templates.Class{
			XFrom: xMin.Add(h.Mul(decimal.NewFromInt32(int32(i)))),
		}
		class.XTo = class.XFrom.Add(h)

		classes = append(classes, &class)
	}

	for i := range classes {
		for _, v := range Variants {
			if v.X.GreaterThanOrEqual(classes[i].XFrom) && v.X.LessThan(classes[i].XTo) {
				classes[i].N++
			}
		}
	}

	c.HTML(http.StatusOK, "classes.tmpl", templates.Classes{
		Classes: classes,
	})
}
