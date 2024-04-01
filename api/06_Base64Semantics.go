package api

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"html/template"
)

const (
	CONDITION = "ZmFsc2U=" // "false" in Base64 encoding
)

func Base64SemanticsTrue(c *gin.Context) {
	// True positive
	if "true" == "true" {
		param := c.Query("param")

		c.HTML(200, "template.html", gin.H{
			"data": template.HTML(param),
		})
	}
}

func Base64SemanticsFalse(c *gin.Context) {
	if base64Decode(CONDITION) == "true" {
		// False positive
		// Analyzers that do not take into account the semantics of standard encoding functions will report a vulnerability
		param := c.Query("param")

		c.HTML(200, "template.html", gin.H{
			"data": template.HTML(param),
		})
	}
}

func base64Decode(s string) string {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	return string(data)
}
