package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var RegisterInput RegisterInput
	c.JSON(http.StatusOK, gin.H{"data": "hello from controller!"})
}
