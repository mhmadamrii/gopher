package main

import (
	"gopher/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)

	r.Run("localhost:5000")
}
