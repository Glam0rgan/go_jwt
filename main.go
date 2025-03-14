package main

import (
	"go_jwt/controllers"
	"go_jwt/middlewares"

	"github.com/gin-gonic/gin"

	"go_jwt/config"
)

func main() {
	r := gin.Default()
	public := r.Group("/api")
	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
	}

	protected := r.Group("/api/admin")
	{
		protected.Use(middlewares.JwtAuthMiddleware())
		protected.GET("/user", controllers.GetInfo)
	}
	config.Init()
	r.Run("0.0.0.0:8000")
}