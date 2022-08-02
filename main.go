package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var gracefulTimeout = time.Second * 10
var shortLatency = time.Millisecond * 200
var longLatency = time.Second * 20

func fast(c *gin.Context) {
	time.Sleep(shortLatency)
	c.JSON(http.StatusOK, gin.H{
		"request": "fast",
	})
}

func slow(c *gin.Context) {
	time.Sleep(longLatency)
	c.JSON(http.StatusOK, gin.H{
		"request": "slow",
	})
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func main() {
	r := gin.Default()

	r.Use(gin.Logger())

	r.GET("/ping", ping)
	r.GET("/fast", fast)
	r.GET("/slow", slow)

	srv := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of gracefulTimeout seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), gracefulTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
