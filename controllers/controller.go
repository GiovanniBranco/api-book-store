package controllers

import (
	"net/http"
	"strconv"

	"github.com/GiovanniBranco/classroom-api/models"
	"github.com/GiovanniBranco/classroom-api/repositories"
	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	cpf := c.Query("cpf")

	if cpf != "" {
		var student models.Student
		repositories.GetStudentByCpf(cpf, &student)

		if student.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Student not found",
			})

			return
		}
		c.JSON(http.StatusOK, student)
		return
	}

	var students []models.Student

	repositories.GetAllStudents(&students)

	c.JSON(http.StatusOK, students)
}

func GetStudentById(c *gin.Context) {

	id := getIdFromParams(c)

	if id == 0 {
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

func GetStudentByCpf(c *gin.Context) {

	cpf := c.Query("cpf")

	var student models.Student

	repositories.GetStudentByCpf(cpf, &student)

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

	if err := student.ValidateStudentData(); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	repositories.CreateStudent(&student)

	c.JSON(http.StatusCreated, student)
}

func UpdateStudent(c *gin.Context) {
	id := getIdFromParams(c)

	if id == 0 {
		return
	}

	var student models.Student
	repositories.GetStudentById(id, &student)

	err := c.ShouldBindBodyWithJSON(&student)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	repositories.UpdateStudent(&student)

	c.Status(http.StatusNoContent)
}

func DeleteStudent(c *gin.Context) {
	id := getIdFromParams(c)

	if id == 0 {
		return
	}

	repositories.DeleteStudent(id)

	c.Status(http.StatusNoContent)
}

func getIdFromParams(c *gin.Context) int {
	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return 0
	}

	return id
}
