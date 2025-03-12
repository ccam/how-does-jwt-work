package Controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gitlab.com/ccam__/how-does-jwt-work/src/models"
)

func GetRoot(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{"title": "hello world"})
}

func GetLogin(c *gin.Context) { // get login page
	c.HTML(200, "login.html", gin.H{"title": "hello world"})
}

func GetAdmin(c *gin.Context) {
	c.HTML(200, "admin.html", gin.H{"title": "hello world"})
}

func GetUserDashboard(c *gin.Context) {
	c.HTML(200, "dashboard.html", gin.H{"title": "hello world"})
}

// login and user management functions
func PostRegister(c *gin.Context) {
	fmt.Println("PostRegister")
}

func PostLogin(c *gin.Context) {
	var user models.User
	c.ShouldBind(&user)

	fmt.Println(user)
}
