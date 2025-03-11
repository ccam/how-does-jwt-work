package Controllers

import (
	"github.com/gin-gonic/gin"
)

func GetRoot(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{"title": "hello world"})
}

func GetLogIn(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{"title": "hello world"})
}

func GetAdmin(c *gin.Context) {
	c.HTML(200, "admin.html", gin.H{"title": "hello world"})
}

func GetUserDashboard(c *gin.Context) {
	c.HTML(200, "dashboad.html", gin.H{"title": "hello world"})
}
