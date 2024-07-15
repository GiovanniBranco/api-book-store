package main

import "github.com/gin-gonic/gin"

func GetAllStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   1,
		"name": "John Doe",
	})
}

func main() {
	r := gin.Default()

	r.GET("/students", GetAllStudents)
	addr := "127.0.0.1:8080/api"
	r.Run(addr)
}
