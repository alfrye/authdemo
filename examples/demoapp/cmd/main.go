package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*.html")
	router.Static("/static", "./internal/public/")
	router.GET("/", getRoot)
	router.GET("/status", getStatus)
	router.POST("/validate", validate)
	auth := router.Group("/bfauth")
	auth.Any("/*Any", authenticate)
	// router.GET("/bfauth/**", authenticate)

	router.Run(":8585")
}

func getRoot(c *gin.Context) {
	c.JSON(http.StatusOK, fmt.Sprintf("%s", time.Now()))
}

func getStatus(c *gin.Context) {
	// c.JSON(http.StatusOK, "ok")
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("status endpoint was called")
	// authheader := c.Request.Header.Get("X-myHeader2")
	// cookieHeader := c.Request.Header.Get("Cookie")
	// c.SetCookie("backend", "jfkjfkjfjeeifjerei", 60*60*24, "", "", true, true)
	c.HTML(http.StatusOK, "status.html", gin.H{
		"content":     "This is a status page...",
		"requestInfo": c.Request.Header,
	})
}

func validate(c *gin.Context) {
	c.JSON(http.StatusOK, fmt.Sprintf("%s", time.Now()))
}

func authenticate(c *gin.Context) {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("authenticate handler fired")

	c.SetCookie("AUTHSESSION", uuid.New().String(), 86400, "/", "", true, true)

	// w.Header().Set("Access-Control-Allow-Origin", validater.Header.Get("Origin"))
	// // w.Header().Set("Access-Control-Allow-Credentials", "true")
	// w.Header().Set("x-myheader", "simplauth")
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("Login succeed"))
	c.JSON(http.StatusOK, "authenticated")
}
