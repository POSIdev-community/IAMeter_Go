package api

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"strings"
)

func CustomFilteringTrue(c *gin.Context) {
	param := c.Query("param")

	// True positive
	c.HTML(200, "template.html", gin.H{
		"data": template.HTML(param),
	})
}

func CustomFilteringFalseFirst(c *gin.Context) {
	param := c.Query("param")

	// False positive
	// Analyzers that do not take into account the semantics of custom filtering functions will report a vulnerability
	c.HTML(200, "template.html", gin.H{
		"data": Filter(param),
	})
}

func CustomFilteringFalseSecond(c *gin.Context) {
	param := c.Query("param")

	// False positive
	// Analyzers that do not take into account the semantics of custom filtering functions will report a vulnerability
	// (CustomFilter.Filter implements the `strings.ReplaceAll("<", "", -1)` logic)
	c.HTML(200, "template.html", gin.H{
		"data": FilterWithStringReplace(param),
	})
}

func Filter(arg string) string {
	var result strings.Builder
	for _, char := range arg {
		if char != '<' && char != '>' && char != '=' {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func FilterWithStringReplace(arg string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(arg, "<", ""), ">", ""), "=", "")
}
