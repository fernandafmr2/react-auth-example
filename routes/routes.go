package routes 

import (
	"github.com/gin-gonic/gin"
	"react-auth-example/controllers"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.GET("/login", controllers.LoginGetHandler())
	g.POST("/login", controllers.LoginPostHandler())
	g.GET("/", controllers.IndexGetHandler())
}

func PrivateRoutes(g *gin.RouterGroup) {
	g.GET("/dashboard", controllers.DashBoardGetHandler())
	g.GET("/logout", controllers.LogoutGetHandler())
}