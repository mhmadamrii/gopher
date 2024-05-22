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

https://seefnasrul.medium.com/create-your-first-go-rest-api-with-jwt-authentication-in-gin-framework-dbe5bda72817
