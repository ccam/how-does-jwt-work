package main

// boilerplate
import (
	"github.com/gin-gonic/gin"
	"gitlab.com/ccam__/how-does-jwt-work/src/Controllers"
	"gitlab.com/ccam__/how-does-jwt-work/src/middleware"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("./templates/*")
	router.Use(Cors())

	// Group routes for admin access and normal user access.

	authRoutes := router.Group("/auth")
	authRoutes.Use(middleware.JWTAuth())
	authRoutes.POST("/register", Controllers.PostRegister)
	authRoutes.POST("/login", Controllers.PostLogin)

	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middleware.JWTAuth())
	adminRoutes.GET("/admin", Controllers.GetAdmin)

	userRoutes := router.Group("/user/")
	userRoutes.Use(middleware.JWTAuth())
	userRoutes.GET("/dashboard", Controllers.GetUserDashboard)

	publicRoutes := router.Group("/public")
	publicRoutes.GET("/", Controllers.GetRoot)
	publicRoutes.GET("/login", Controllers.GetLogin)
	publicRoutes.POST("/user-login", Controllers.PostLogin)

	router.Run("localhost:8000")
}
