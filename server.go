package main

import (
	"github.com/gin-gonic/gin"
	"react-auth-example/controllers"
	"react-auth-example/middleware"
	"react-auth-example/globals"
	"github.com/gin-contrib/sessions/cookie"
	"react-auth-example/routes"
	"github.com/gin-contrib/sessions"
	"net/http"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8000")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	userRepo := controllers.New()
	r.POST("/users", userRepo.CreateUser)
	r.GET("/users", userRepo.GetUsers)
	r.GET("/users/:id", userRepo.GetUser)
	r.PUT("/users/:id", userRepo.UpdateUser)
	r.DELETE("/users/:id", userRepo.DeleteUser)

	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*.html")

	r.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := r.Group("/")
	routes.PublicRoutes(public)

	private := r.Group("/")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private)

	return r
}