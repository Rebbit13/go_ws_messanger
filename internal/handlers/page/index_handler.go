package page

import (
	"github.com/gin-gonic/gin"
)

type PageHandler struct{}

func (handler *PageHandler) getIndexPage(context *gin.Context) {
	context.Header("content-type", "text/html")
	context.HTML(200, "index.html", gin.H{
		"title": "Main Page",
	})
}

func BindHandler(router *gin.Engine) {
	router.Static("/static", "./web/templates/static")
	router.LoadHTMLGlob("web/templates/*.html")
	var handler = PageHandler{}
	router.GET("/", handler.getIndexPage)
}
