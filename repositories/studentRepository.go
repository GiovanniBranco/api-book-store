package repositories

import (
	"github.com/GiovanniBranco/classroom-api/database"
	"github.com/GiovanniBranco/classroom-api/models"
)

func GetAllStudents(students *[]models.Student) {
	database.DB.Find(students)
}

func GetStudentById(id int, student *models.Student) {
	database.DB.Find(student, id)
}

func CreateStudent(student *models.Student) {
	database.DB.Create(student)
}
