package factory

import (
	"student-service/database"
	"student-service/internal/repository"
)

type Factory struct {
	StudentRepository repository.Student
	MajorRepository   repository.Major
	ClassRepository   repository.Class
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		repository.NewStudentRepository(db),
		repository.NewMajorRepository(db),
		repository.NewClassRepository(db),
	}
}
