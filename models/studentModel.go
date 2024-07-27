package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string `json:"name" validate:"nonzero"`
	Cpf  string `json:"cpf" validate:"len=9, regexp=^[0-9]*$"`
}

func (student *Student) ValidateStudentData() error {
	err := validator.Validate(student)
	if err != nil {
		return err
	}

	return nil
}
