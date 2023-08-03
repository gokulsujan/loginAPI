package main

import (
	"github.com/gin-gonic/gin"
)

var LoginData struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(c *gin.Context) {
	if err := c.ShouldBindJSON(&LoginData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	const validUserName = "superadmin"
	const validPassword = "superpassword"

	if LoginData.UserName == validUserName && LoginData.Password == validPassword {
		c.JSON(200, gin.H{"message": "Login successful"})
	} else {
		c.JSON(401, gin.H{"message": "Invalid Credentials"})
	}
}

func main() {
	r := gin.Default()
	r.POST("/login", loginHandler)
	r.Run(":5050")
}
