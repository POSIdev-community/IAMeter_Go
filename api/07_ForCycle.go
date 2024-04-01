package api

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

func ForCycleTrue(c *gin.Context) {
	sum := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 15; j++ {
			sum += i + j
		}
	}

	if sum == 1725 {
		param := c.Query("param")

		// True positive
		c.HTML(200, "template.html", gin.H{
			"data": template.HTML(param),
		})
	}
}

func ForCycleFalseFirst(c *gin.Context) {
	sum := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 15; j++ {
			sum += i + j
		}
	}

	if sum != 1725 {
		param := c.Query("param")

		// False positive
		// Analyzers that approximate or ignore the cycle interpretation will report a vulnerability
		c.HTML(200, "template.html", gin.H{
			"data": template.HTML(param),
		})
	}
}

func ForCycleFalseSecond(c *gin.Context) {
	sum := 0
	for i := range [10]int{} {
		for j := range make([]int, 15) {
			sum += i + j
		}
	}

	if sum != 1725 {
		param := c.Query("param")

		// False positive
		// Analyzers that approximate or ignore the cycle interpretation will report a vulnerability
		c.HTML(200, "template.html", gin.H{
			"data": template.HTML(param),
		})
	}
}
