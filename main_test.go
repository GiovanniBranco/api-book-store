package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GiovanniBranco/classroom-api/controllers"
	"github.com/GiovanniBranco/classroom-api/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var StudentId int

func getRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	routes.GET("/api/students", controllers.GetAllStudents)
	routes.GET("/api/students/:id", controllers.GetStudentById)

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

	expectResponse := `
  {
    "ID": 1,
    "CreatedAt": "2024-07-27T15:28:25.214666-03:00",
    "UpdatedAt": "2024-07-27T15:28:25.214666-03:00",
    "DeletedAt": null,
    "name": "John Doe",
    "cpf": "12345678909"
  }`
	routes := getRoutes()

	req, _ := http.NewRequest("GET", "/api/students/1", nil)
	res := httptest.NewRecorder()

	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.JSONEq(t, expectResponse, res.Body.String())
}
