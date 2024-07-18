package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/GiovanniBranco/classroom-api/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/api/students", controllers.GetAllStudents)
	r.GET("/api/students/:id", controllers.GetStudentById)
	r.POST("/api/students", controllers.CreateStudent)
	addr := "127.0.0.1:8080"
	r.Run(addr)
}
