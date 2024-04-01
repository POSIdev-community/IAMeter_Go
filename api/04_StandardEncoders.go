package api

import (
	"github.com/gin-gonic/gin"
	"html"
	"html/template"
)

func StandardEncodersTrue(c *gin.Context) {
	param := c.Query("param")

	c.HTML(200, "template.html", gin.H{
		// True positive
		"data": template.HTML(param),
	})
}

func StandardEncodersFalse(c *gin.Context) {
	param := c.Query("param")

	c.HTML(200, "template.html", gin.H{
		// An analyzer that ignores the semantics of standard filter functions will report a vulnerability here
		// False positive
		"data": html.EscapeString(param),
	})
}
