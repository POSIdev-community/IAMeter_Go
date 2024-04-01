package main

import (
	"github.com/POSIdev-community/iameter-go/api"
	"github.com/gin-gonic/gin"
	"strings"
)

type Link struct {
	Url            string
	Classification string
}

type RouteFile struct {
	Name  string
	Links []Link
}

const (
	TAINTED_PARAM = "?param=<script>alert(decodeURIComponent('0'))</script>"
)

func main() {
	router := gin.New()
	router.LoadHTMLGlob("templates/*")
	initRoutes(router)
	router.Run(":8080")
}

func initRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", buildRouteFiles(router))
	})

	router.GET("functional-flow/true", api.FunctionalFlowTrue)
	router.GET("functional-flow/false", api.FunctionalFlowFalse)
	router.GET("broken-code/false", api.BrokenCodeFalse)
	router.GET("unreachable-code/true", api.UnreacheableCodeTrue)
	router.GET("unreachable-code/false", api.UnreacheableCodeFalse)
	router.GET("standard-encoders/true", api.StandardEncodersTrue)
	router.GET("standard-encoders/false", api.StandardEncodersFalse)
	router.GET("custom-filtering/true", api.CustomFilteringTrue)
	router.GET("custom-filtering/false/1", api.CustomFilteringFalseFirst)
	router.GET("custom-filtering/false/2", api.CustomFilteringFalseSecond)
	router.GET("base64-semantics/true", api.Base64SemanticsTrue)
	router.GET("base64-semantics/false", api.Base64SemanticsFalse)
	router.GET("for-cycle/true", api.ForCycleTrue)
	router.GET("for-cycle/false/1", api.ForCycleFalseFirst)
	router.GET("for-cycle/false/2", api.ForCycleFalseSecond)
	router.GET("string-semantics/true", api.StringSemanticsTrue)
	router.GET("string-semantics/false/1", api.StringSemanticsFalseFirst)
	router.GET("string-semantics/false/2", api.StringSemanticsFalseSecond)
	router.GET("string-semantics/false/3", api.StringSemanticsFalseThird)
	router.GET("type-cast/true", api.TypeCastTrue)
	router.GET("type-cast/false", api.TypeCastFalse)
}

func buildRouteFiles(router *gin.Engine) []RouteFile {
	routes := router.Routes()
	routeFiles := make([]RouteFile, 0)
	fileMap := make(map[string][]Link)

	for _, route := range routes {
		if route.Path != "/" {
			fileName,
				classification := parsePath(route.Path)
			fileMap[fileName] = append(fileMap[fileName], Link{
				Url:            route.Path + TAINTED_PARAM,
				Classification: classification,
			})
		}
	}

	for fileName, links := range fileMap {
		routeFiles = append(routeFiles, RouteFile{
			Name:  fileName,
			Links: links,
		})
	}

	return routeFiles
}

func parsePath(path string) (string, string) {
	parts := strings.Split(path, "/")
	fileName := strings.Title(strings.ReplaceAll(parts[1], "-", " "))
	classification := strings.Title(parts[2])
	return fileName, classification
}
