package controller

import (
    "github.com/gin-gonic/gin"
)


func GetRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("view/template/*.html")

	r.GET("/", indexPage)
	r.POST("/threads", createThread)
	r.GET("/threads/:ThreadId", getThreadPage)
	r.POST("/threads/:ThreadId/responses", postResponse)                     

	return r
}



