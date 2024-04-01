package api

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

func FunctionalFlowTrue(c *gin.Context) {
	param := c.Query("param")

	// True positive
	c.HTML(200, "template.html", gin.H{
		"data": template.HTML(param),
	})
}

func FunctionalFlowFalse(c *gin.Context) {
	param := c.Query("param")

	pvo := func(arg string) string {
		return arg
	}

	// False negative
	// Analyzers that do not interpret the execution flow based on functional data flows will NOT report a vulnerability
	c.HTML(200, "template.html", gin.H{"data": template.HTML(pvo(param))})
}
