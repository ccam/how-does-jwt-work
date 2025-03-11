package main

// boilerplate
import (
	"github.com/gin-gonic/gin"
	"gitlab.com/ccam__/how-does-jwt-work/src/Controllers"
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
	router.GET("/", Controllers.GetRoot)
	router.GET("/admin/", Controllers.GetAdmin)
	router.GET("/dashboard/", Controllers.GetUserDashboard)
	router.GET("/login/", Controllers.GetLogIn)
	router.Run("localhost:8080")
}
