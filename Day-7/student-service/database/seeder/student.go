package seeder

import (
	"log"
	"time"

	"student-service/internal/model"

	"gorm.io/gorm"
)

func studentSeeder(db *gorm.DB) {
	now := time.Now()
	var students = []model.Student{
		{
			Fullname: "Vincent L. Hubbard",
			Email:    "vincentlhubbard@edu.ac.id",
			Password: "$2a$10$rfpS/jJ.a5J9seBM5sNPTeMQ0iVcAjoox3TDZqLE7omptkVQfaRwW", // 123abcABC!
			ClassID:  1,
			MajorID:  1,
			Common:   model.Common{ID: 1, CreatedAt: now, UpdatedAt: now},
		},
		{
			Fullname: "Devon C. Thomas",
			Email:    "devoncthomas@edu.ac.id",
			Password: "$2a$10$rfpS/jJ.a5J9seBM5sNPTeMQ0iVcAjoox3TDZqLE7omptkVQfaRwW", // 123abcABC!
			ClassID:  2,
			MajorID:  1,
			Common:   model.Common{ID: 2, CreatedAt: now, UpdatedAt: now},
		},
		{
			Fullname: "Bettina M. Easter",
			Email:    "bettinameaster@edu.ac.id",
			Password: "$2a$10$rfpS/jJ.a5J9seBM5sNPTeMQ0iVcAjoox3TDZqLE7omptkVQfaRwW", // 123abcABC!
			ClassID:  2,
			MajorID:  2,
			Common:   model.Common{ID: 3, CreatedAt: now, UpdatedAt: now},
		},
	}
	if err := db.Create(&students).Error; err != nil {
		log.Printf("cannot seed data students, with error %v\n", err)
	}
	log.Println("success seed data students")
}
