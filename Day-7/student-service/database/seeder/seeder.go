package seeder

import (
	"student-service/database"

	"gorm.io/gorm"
)

type seed struct {
	DB *gorm.DB
}

func NewSeeder() *seed {
	return &seed{database.GetConnection()}
}

func (s *seed) SeedAll() {
	classSeeder(s.DB)
	majorSeeder(s.DB)
	studentSeeder(s.DB)
}

func (s *seed) DeleteAll() {
	s.DB.Exec("DELETE FROM students")
	s.DB.Exec("DELETE FROM majors")
	s.DB.Exec("DELETE FROM classes")
}
