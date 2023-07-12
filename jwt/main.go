package main

import (
	"github.com/gin-gonic/gin"
	"github.com/naqet/learning-go/jwt/controllers"
	"github.com/naqet/learning-go/jwt/initializers"
	"github.com/naqet/learning-go/jwt/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
	initializers.SyncDb()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run() // listen and serve on 0.0.0.0:8080
}
