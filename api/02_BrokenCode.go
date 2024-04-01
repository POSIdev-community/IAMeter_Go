package api

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

func BrokenCodeFalse(c *gin.Context) {
	param := c.Query("param")

	// True positive
	c.HTML(200, "template.html", gin.H{
		"data": template.HTML(param),
	})

	// For analyzers that require compilable code, delete or comment out this fragment
	/*
		argument := "harmless value"

		UnknownType.Property1 = output
		UnknownType.Property2 = UnknownType.Property1
		UnknownType.Property3 = cond1

		if UnknownType.Property3 == nil {
			argument = UnknownType.Property2
		}

		// An analyzer that ignores noncompiled code will report a vulnerability here
		// False positive
		c.HTML(200, "template.html", gin.H{
			"data": "Hello " + argument,
		})*/
}
