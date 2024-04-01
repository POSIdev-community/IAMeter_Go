package api

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"strconv"
)

func TypeCastTrue(c *gin.Context) {
	param := c.Query("param")

	// True positive
	c.HTML(200, "template.html", gin.H{
		"data": template.HTML(param),
	})
}

func TypeCastFalse(c *gin.Context) {
	param, err := strconv.Atoi(c.Query("param"))

	if err != nil {
		param = 0
	}

	// False positive
	// Analyzers that do not take into account that the parameter is converted to a number will report a vulnerability
	c.HTML(200, "template.html", gin.H{
		"param": param,
	})
}
