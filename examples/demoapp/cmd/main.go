package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("./templates/*.html")
	router.Static("/static", "./internal/public/")
	router.GET("/", getRoot)
	router.GET("/status", getStatus)
	router.POST("/validate", validate)

	router.Run(":8585")

}

func getRoot(c *gin.Context) {

	c.JSON(http.StatusOK, fmt.Sprintf("%s", time.Now()))
}

func getStatus(c *gin.Context) {

	//c.JSON(http.StatusOK, "ok")

	c.SetCookie("backend", "jfkjfkjfjeeifjerei", 60*60*24, "", "", true, true)
	c.HTML(http.StatusOK, "status.html", gin.H{
		"content": "This is a status page...",
	})
}

func validate(c *gin.Context) {
	c.JSON(http.StatusOK, fmt.Sprintf("%s", time.Now()))
}
