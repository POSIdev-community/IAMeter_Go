package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"strings"
)

const (
	COND = "ZmFsc2U=" // "false" in Base64 encoding
)

func StringSemanticsTrue(c *gin.Context) {
	condTrue := "true"
	if condTrue == "true" {
		param := c.Query("param")

		// True positive
		c.HTML(200, "template.html", gin.H{"data": template.HTML(param)})
	}
}

func StringSemanticsFalseFirst(c *gin.Context) {
	param := c.Query("param")
	if COND+"true" == "true" {
		// An analyzer that does not interpret the semantics of standard library types will report a vulnerability here
		// False positive
		c.HTML(200, "template.html", gin.H{"data": template.HTML(param)})
	} else {
		c.String(200, "Hello")
	}
}

func StringSemanticsFalseSecond(c *gin.Context) {
	param := c.Query("param")
	if strings.Join([]string{COND, "true"}, "") == "true" {
		// An analyzer that does not interpret the semantics of standard library types will report a vulnerability here
		// False positive
		c.HTML(200, "template.html", gin.H{"data": template.HTML(param)})
	} else {
		c.String(200, "Hello")
	}
}

func StringSemanticsFalseThird(c *gin.Context) {
	param := c.Query("param")

	if fmt.Sprintf("%s true", COND) == "true" {
		// An analyzer that does not interpret the semantics of standard library types will report a vulnerability here
		// False positive
		c.HTML(200, "template.html", gin.H{"data": template.HTML(param)})
	} else {
		c.String(200, "Hello")
	}
}
