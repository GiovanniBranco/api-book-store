package controllers

import (
	"net/http"
	"strconv"

	"github.com/GiovanniBranco/classroom-api/models"
	"github.com/GiovanniBranco/classroom-api/repositories"
	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	var students []models.Student

	repositories.GetAllStudents(&students)

	c.JSON(http.StatusOK, students)
}

func GetStudentById(c *gin.Context) {

	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	var student models.Student

	repositories.GetStudentById(id, &student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})

		return
	}

	c.JSON(http.StatusOK, student)

}

func CreateStudent(c *gin.Context) {
	var student models.Student

	err := c.ShouldBindBodyWithJSON(&student)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	repositories.CreateStudent(&student)

	c.JSON(http.StatusCreated, student)
}
