package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/ping-1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong-2",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello page",
		})
	})

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Test page",
		})
	})

	r.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Test 2 page",
		})
	})

	r.GET("/test3", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Test 3 page",
		})
	})
	r.GET("/test4", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Test 4 page",
		})
	})

	r.GET("/test5", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Test 5 page",
		})
	})

	r.Run()
}
