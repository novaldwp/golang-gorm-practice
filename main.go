package main

import (
	"golang_gorm_practice/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default() // router

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	userRepo := controllers.New()
	r.POST("/users", userRepo.CreateUser)
	r.GET("/users", userRepo.GetUsers)
	r.GET("/users/:id", userRepo.GetUserById)
	r.PUT("/users/:id", userRepo.UpdateUser)
	r.DELETE("users/:id", userRepo.DeleteUser)

	return r
}
