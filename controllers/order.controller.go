package controllers

import (
	"net/http"
	"chatting-api/controllers"
	"github.com/gin-gonic/gin"
)

func Getting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "GET"})
}

func Posting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "POST"})
}

func Deleting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "DELETE"})
}

func Patching(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "PATCH"})
}

func Options(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "OPTIONS"})
}
