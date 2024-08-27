package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./static/*.html")
	r.Static("/css", "./css")
	r.Static("/img", "./img")
	r.StaticFS("/.git", http.Dir("./.git"))

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run(":2000")
}
