package repositories

import (
	"github.com/GiovanniBranco/classroom-api/database"
	"github.com/GiovanniBranco/classroom-api/models"
)

func GetAllStudents(students *[]models.Student) {
	database.DB.Find(students)
}

func GetStudentById(id int, student *models.Student) {
	database.DB.First(student, id)
}

func GetStudentByCpf(cpf string, student *models.Student) {
	database.DB.Where(&models.Student{
		Cpf: cpf,
	}).First(student)
}

func CreateStudent(student *models.Student) {
	database.DB.Create(student)
}

func UpdateStudent(student *models.Student) {
	database.DB.Model(&student).UpdateColumns(student)
}

func DeleteStudent(id int) {
	database.DB.Delete(&models.Student{}, id)
}
