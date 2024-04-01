package api

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

func UnreacheableCodeTrue(c *gin.Context) {
	param := c.Query("param")

	c.HTML(200, "template.html", gin.H{
		// True positive
		"data": template.HTML(param),
	})
}

func UnreacheableCodeFalse(c *gin.Context) {
	cond1 := "ZmFsc2U=" // "false" in Base64 encoding

	if cond1 == "" {
		param := c.Query("param")

		// False positive
		// Analyzers that do not take into account reachability of execution paths will report a vulnerability
		c.HTML(200, "template.html", gin.H{
			// True positive
			"data": template.HTML(param),
		})
	} else {
		c.HTML(200, "template.html", gin.H{
			"data": "Hello",
		})
	}
}
