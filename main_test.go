package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GiovanniBranco/classroom-api/controllers"
	"github.com/GiovanniBranco/classroom-api/database"
	"github.com/GiovanniBranco/classroom-api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var StudentId int

func getRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	routes.GET("/api/students", controllers.GetAllStudents)
	routes.GET("/api/students/:id", controllers.GetStudentById)
	routes.PUT("/api/students/:id", controllers.UpdateStudent)

	return routes
}

func connectDatabase() {
	database.ConnectDatabase()
}

func TestGetAllStudents(t *testing.T) {
	connectDatabase()
	expectResponse := `[
  {
    "ID": 1,
    "CreatedAt": "2024-07-27T15:28:25.214666-03:00",
    "UpdatedAt": "2024-07-27T15:28:25.214666-03:00",
    "DeletedAt": null,
    "name": "John Doe",
    "cpf": "12345678909"
  },
  {
    "ID": 2,
    "CreatedAt": "2024-07-27T17:14:51.303646-03:00",
    "UpdatedAt": "2024-07-27T17:14:51.303646-03:00",
    "DeletedAt": null,
    "name": "",
    "cpf": "12345678909"
  },
  {
    "ID": 3,
    "CreatedAt": "2024-07-27T17:16:04.856538-03:00",
    "UpdatedAt": "2024-07-27T17:16:04.856538-03:00",
    "DeletedAt": null,
    "name": "John Doe",
    "cpf": "12345678909"
  }
]`
	routes := getRoutes()

	req, _ := http.NewRequest("GET", "/api/students", nil)
	res := httptest.NewRecorder()

	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.JSONEq(t, expectResponse, res.Body.String())
}

func TestGetStudentById(t *testing.T) {
	connectDatabase()

	var studentMock models.Student

	routes := getRoutes()

	req, _ := http.NewRequest("GET", "/api/students/1", nil)
	res := httptest.NewRecorder()

	routes.ServeHTTP(res, req)

	json.Unmarshal(res.Body.Bytes(), &studentMock)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "John Doe", studentMock.Name)
	assert.Equal(t, "12345678909", studentMock.Cpf)
}

func TestUpdateStudent(t *testing.T) {
	connectDatabase()

	routes := getRoutes()

	studentUpdated := models.Student{Name: "Nome do Aluno Teste", Cpf: "47123456789"}

	payload, _ := json.Marshal(studentUpdated)

	req, _ := http.NewRequest("PUT", "/api/students/1", bytes.NewBuffer(payload))
	res := httptest.NewRecorder()

	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusNoContent, res.Code)
}
